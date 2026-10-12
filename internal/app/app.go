package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Vafeda/plc-historian/internal/config"
	"github.com/Vafeda/plc-historian/internal/database"
	"github.com/Vafeda/plc-historian/internal/deadband"
	"github.com/Vafeda/plc-historian/internal/services"
	"github.com/Vafeda/plc-historian/internal/transport"

	"github.com/gopcua/opcua/debug"
)

func Run() {
	// Чтение конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	debug.Enable = cfg.Debug
	log.SetFlags(log.Ltime | log.Lmicroseconds)

	// Стартовый контекст
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Запуск и отложенное закрытие пула соединения с БД
	pool, err := database.Open(ctx, cfg.DSN)
	if err != nil {
		log.Fatalf("база данных: %v", err)
	}
	defer database.Close(pool)

	history := services.NewHistory(database.NewHistory(pool))

	filter := deadband.New(0.5, nil)

	var wg sync.WaitGroup
	// Запуск батчера, который собирает Samples
	wg.Go(func() {
		history.RunBatcher(ctx, filter)
	})

	transport.Run(ctx, history, cfg.Endpoint)

	wg.Wait()
}
