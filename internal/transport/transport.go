package transport

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Vafeda/plc-historian/internal/models"
	"github.com/Vafeda/plc-historian/internal/services"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

// Проверка живости соединения.
//
// Раньше watchdog сбрасывался только сообщениями из notifyCh — но gopcua
// НЕ кладёт keep-alive в notifyCh, туда приходят только реальные изменения
// данных. Пока кран стоит и значения не меняются, watchdog через 10с
// убивал вполне живое соединение. Теперь живость проверяем отдельным
// лёгким Read (текущее время сервера) раз в heartbeatInterval.
const (
	heartbeatInterval = 3 * time.Second
	heartbeatTimeout  = 3 * time.Second
	// Если за это время не было ни одного успешного heartbeat (и ни одного
	// изменения данных) — считаем соединение мёртвым и переподключаемся.
	// Нужно на случай, если Read на полуоткрытом сокете повиснет навсегда.
	livenessTimeout = 10 * time.Second
)

// Сервер отвечает StatusBadTooManyOperations, если запросить подписку сразу
// на все ноды одним CreateMonitoredItemsRequest — разбиваем на пачки.
const monitorBatchSize = 50

// Серверный deadband не используем: OPC UA сервер CODESYS 3.5.17.30 на этом
// ПЛК отклоняет любой DataChangeFilter с deadband (Absolute и Percent) —
// StatusBadMonitoredItemFilterUnsupported. Фильтровать нужно у себя.

// ns=0;i=2258 — Server_ServerStatus_CurrentTime, стандартный узел любого
// OPC UA сервера. Читать его дёшево, подходит для heartbeat.
var serverCurrentTimeNode = ua.NewNumericNodeID(0, id.Server_ServerStatus_CurrentTime)

func ptr[T any](v T) *T { return &v }

func Run(ctx context.Context, history *services.History, endpoint string) {
	backoff := time.Second
	for ctx.Err() == nil {
		err := runSession(ctx, history, endpoint)
		if err != nil {
			log.Printf("session ended: %v — reconnecting in %s", err, backoff)
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(backoff):
		}
		// простой capped backoff: 1с -> 2с -> 4с -> ... -> максимум 10с
		backoff = min(backoff*2, 10*time.Second)
	}
}

// runSession живёт, пока соединение живо. Как только что-то ломается —
// возвращает ошибку, и main() пересоздаёт всё с нуля.
func runSession(ctx context.Context, history *services.History, endpoint string) error {
	// Контекст сессии: отменяется при выходе из runSession, чтобы
	// горутина heartbeat гарантированно завершилась.
	sessCtx, sessCancel := context.WithCancel(ctx)
	defer sessCancel()

	endpoints, err := opcua.GetEndpoints(sessCtx, endpoint)
	if err != nil {
		return err
	}
	fmt.Println(*endpoints[0], *endpoints[1])

	ep, err := opcua.SelectEndpoint(endpoints, "None", ua.MessageSecurityModeNone)
	if err != nil {
		return err
	}
	ep.EndpointURL = endpoint
	fmt.Println(*ep)

	opts := []opcua.Option{
		opcua.SecurityPolicy("None"),
		opcua.SecurityModeString("None"),
		opcua.AuthAnonymous(),
		opcua.SecurityFromEndpoint(ep, ua.UserTokenTypeAnonymous),
		opcua.AutoReconnect(false),
	}

	c, err := opcua.NewClient(ep.EndpointURL, opts...)
	if err != nil {
		return err
	}
	if err := c.Connect(sessCtx); err != nil {
		return err
	}
	// Закрываем клиента в фоне и не ждём результата: на мёртвом соединении
	// запись в разорванный сокет может зависнуть дольше любого таймаута.
	defer func() {
		go func() {
			closeCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			c.Close(closeCtx)
		}()
	}()

	// Буфер — чтобы библиотека не блокировалась, пока мы ещё создаём
	// monitored items (читать канал начинаем только после этого).
	notifyCh := make(chan *opcua.PublishNotificationData, 256)
	sub, err := c.Subscribe(sessCtx, &opcua.SubscriptionParameters{
		Interval:          100 * time.Millisecond,
		MaxKeepAliveCount: 3,   // keep-alive примерно раз в 300мс простоя
		LifetimeCount:     300, // подписка живёт минимум 30с без publish'ей
	}, notifyCh)
	if err != nil {
		return err
	}
	defer func() {
		go func() {
			cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			sub.Cancel(cancelCtx)
		}()
	}()

	reqs := make([]*ua.MonitoredItemCreateRequest, 0, len(nodes))
	validNodes := make([]string, 0, len(nodes))
	for _, node := range nodes {
		n, err := ua.ParseNodeID(node)
		if err != nil {
			log.Printf("skip %s: %v", node, err)
			continue
		}
		// ClientHandle = items.id тега в БД (см. handleToName в nodes.go):
		// сервер возвращает его в каждом уведомлении, и он же пишется в history.itemid.
		req := opcua.NewMonitoredItemCreateRequestWithDefaults(n, ua.AttributeIDValue, nodeHandle[node])
		// По умолчанию SamplingInterval=0 ("как можно быстрее") — явно 100мс.
		req.RequestedParameters.SamplingInterval = 100
		reqs = append(reqs, req)
		validNodes = append(validNodes, node)
	}

	ok := 0
	for start := 0; start < len(reqs); start += monitorBatchSize {
		end := min(start+monitorBatchSize, len(reqs))
		res, err := sub.Monitor(sessCtx, ua.TimestampsToReturnBoth, reqs[start:end]...)
		if err != nil {
			return err
		}
		for i, r := range res.Results {
			if r.StatusCode == ua.StatusOK {
				ok++
			} else {
				log.Printf("failed to monitor %s: %v", validNodes[start+i], r.StatusCode)
			}
		}
	}
	log.Printf("subscription %v created: %d/%d tags ok", sub.SubscriptionID, ok, len(reqs))

	// Heartbeat в отдельной горутине: сообщает об успехе в alive,
	// об ошибке — в hbErr.
	alive := make(chan struct{}, 1)
	hbErr := make(chan error, 1)
	go heartbeat(sessCtx, c, alive, hbErr)

	liveness := time.NewTimer(livenessTimeout)
	defer liveness.Stop()
	resetLiveness := func() {
		if !liveness.Stop() {
			select {
			case <-liveness.C:
			default:
			}
		}
		liveness.Reset(livenessTimeout)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-hbErr:
			return fmt.Errorf("heartbeat failed: %w", err)
		case <-alive:
			resetLiveness()
		case <-liveness.C:
			return errors.New("liveness: no successful heartbeat for " + livenessTimeout.String())
		case msg := <-notifyCh:
			resetLiveness() // данные пришли — соединение точно живо

			if msg.Error != nil {
				return msg.Error // выходим -> main() переподключится
			}
			switch x := msg.Value.(type) {
			case *ua.DataChangeNotification:
				for _, item := range x.MonitoredItems {
					name := handleToName[item.ClientHandle]
					dv := item.Value // значение + статус + метки времени
					if dv == nil {
						log.Printf("%s: пустое уведомление", name)
						continue
					}
					// При плохом статусе сервер может прислать DataValue без
					// самого значения (dv.Value == nil) — Value() на нём паникует.
					if dv.Status != ua.StatusOK || dv.Value == nil {
						log.Printf("%s: плохое значение, статус %v", name, dv.Status)
						continue
					}

					ts := dv.SourceTimestamp
					if ts.IsZero() {
						ts = time.Now().UTC() // сервер не прислал метку — берём своё время
					}

					sample := models.Sample{
						ItemId:  item.ClientHandle,
						ItemTS:  ts,
						RegTS:   time.Now().UTC(),
						Quality: dv.Status,
					}

					// TODO: Вместо ptr можно использовать встроенный new, но отображается как ошибка в IDE при этом спокойно компилируется
					switch dv.Value.Type() {
					case ua.TypeIDBoolean:
						sample.ValBool = ptr(dv.Value.Bool())
					case ua.TypeIDSByte, ua.TypeIDInt16, ua.TypeIDInt32, ua.TypeIDInt64:
						sample.ValInt = ptr(dv.Value.Int())
					case ua.TypeIDByte, ua.TypeIDUint16, ua.TypeIDUint32:
						sample.ValInt = ptr(int64(dv.Value.Uint()))
					case ua.TypeIDFloat, ua.TypeIDDouble:
						sample.ValDouble = ptr(dv.Value.Float())
					case ua.TypeIDString:
						sample.ValStr = ptr(dv.Value.String())
					default:
						log.Printf("%s: тип %v не поддерживается", name, dv.Value.Type())
						continue
					}

					history.Push(sample)
				}
			default:
				log.Printf("unexpected type %T", msg.Value)
			}
		}
	}
}

// heartbeat раз в heartbeatInterval читает текущее время сервера.
// Успех — сигнал в alive, ошибка — в hbErr и выход.
func heartbeat(ctx context.Context, c *opcua.Client, alive chan<- struct{}, hbErr chan<- error) {
	t := time.NewTicker(heartbeatInterval)
	defer t.Stop()
	req := &ua.ReadRequest{
		NodesToRead: []*ua.ReadValueID{
			{NodeID: serverCurrentTimeNode, AttributeID: ua.AttributeIDValue},
		},
		TimestampsToReturn: ua.TimestampsToReturnNeither,
	}
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			rctx, cancel := context.WithTimeout(ctx, heartbeatTimeout)
			_, err := c.Read(rctx, req)
			cancel()
			if err != nil {
				if ctx.Err() != nil {
					return // сессию и так закрывают
				}
				select {
				case hbErr <- err:
				default:
				}
				return
			}
			select {
			case alive <- struct{}{}:
			default:
			}
		}
	}
}
