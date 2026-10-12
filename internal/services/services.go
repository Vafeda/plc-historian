package services

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"github.com/Vafeda/plc-historian/internal/database"
	"github.com/Vafeda/plc-historian/internal/deadband"
	"github.com/Vafeda/plc-historian/internal/models"
)

const (
	// Раз в столько времени накопленная пачка выводится (потом — пишется в БД).
	batchInterval = time.Second
	// Размер канала между приёмом данных с ПЛК и сборщиком пачек.
	samplesChanSize = 10000
)

type History struct {
	db             *database.History
	samples        chan models.Sample
	droppedSamples atomic.Uint64
}

func NewHistory(db *database.History) *History {
	return &History{
		db:      db,
		samples: make(chan models.Sample, samplesChanSize),
	}
}

func (h *History) Push(sample models.Sample) {
	select {
	case h.samples <- sample:
	default:
		h.droppedSamples.Add(1)
	}
}

// runBatcher копит значения в пачку и раз в batchInterval выводит её.
// Работает всё время программы и переживает переподключения к ПЛК.
func (h *History) RunBatcher(ctx context.Context, filter *deadband.Filter) {
	t := time.NewTicker(batchInterval)
	defer t.Stop()

	var raw, arch []models.Sample

	flush := func(ctx context.Context) {
		if d := h.droppedSamples.Swap(0); d > 0 {
			log.Printf("ВНИМАНИЕ: выброшено %d значений — сборщик пачек не успевал", d)
		}

		// Пришедший батч нулевой. Нет смысла его отдавать в слой базы данных
		if len(raw) == 0 {
			return
		}

		if err := h.db.Write(ctx, raw, arch); err != nil {
			filter.Reset()
			log.Printf("error: %v, count raw data lost : %d, count arch data lost : %d", err, len(raw), len(arch))
		}

		raw = raw[:0] // память среза переиспользуется для следующей пачки
		arch = arch[:0]
	}

	for {
		select {
		case s := <-h.samples:
			raw = append(raw, s)
			v, _ := filter.Pass(s)
			if v {
				arch = append(arch, s)
			}
		case <-t.C:

			flush(ctx)
		case <-ctx.Done():
			// забрать то, что уже лежит в канале, и вывести напоследок
			for {
				select {
				case s := <-h.samples:
					raw = append(raw, s)
					v, _ := filter.Pass(s)
					if v {
						arch = append(arch, s)
					}
				default:
					ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
					flush(ctx)
					cancel()
					return
				}
			}
		}
	}
}
