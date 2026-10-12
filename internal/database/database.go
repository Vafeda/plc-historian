package database

import (
	"context"
	//"errors"
	"fmt"
	//"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	//"os"
	"time"
	//"github.com/Vafeda/wallet/internal/config"
)

func Open(ctx context.Context, dsn string /*log *slog.Logger*/) (*pgxpool.Pool, error) {
	//connStr := fmt.Sprintf(
	//	"postgres://%s:%s@%s:%s/%s?sslmode=disable",
	//	cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	//)

	// TODO: Логирование результата парсинга конфигурации ссылки. INFO при удаче/ERROR при неудачной обработке (надо искать решение)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	// TODO: Логирования результата выставления всех настроек для подключения к БД
	// Формирование конфигурации подключения к БД. Можно через строку, либо через код. В будущем легче думаю через код это делать.
	//config.MaxConns = int32(cfg.MaxConns)
	//config.MinConns = int32(cfg.MinConns)
	//config.MaxConnLifetime = cfg.MaxConnLifetime
	//config.MaxConnIdleTime = cfg.MaxConnIdleTime

	// TODO: Информирование о создание подключения или нет
	// Создание пула для соединения с базой
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// TODO: А что если нам нужна миграция, но при этом база не пинганулась?
	// TODO: От чьего лица запускаеться подключение к базе данных? Мы делаем нового пользователя или используем стокового?
	// TODO: При использование стокового есть вероятность дропнуть таблицу самолично
	err = pool.Ping(ctxPing)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// TODO: Миграция для будущего.
	//if cfg.NeedMigration {
	//	err = runMigration(pool, cfg)
	//	if err != nil {
	//		var pgErr *pgconn.PgError
	//		if errors.As(err, &pgErr) {
	//			if pgErr.Code != "42P07" {
	//				return nil, err
	//			}
	//
	//			fmt.Fprintf(os.Stderr, "Migration failed (continuing anyway): %v\n", err)
	//		}
	//	}
	//}

	fmt.Println("Database connection pool created successfully")
	return pool, nil
}

func Close(conn *pgxpool.Pool) {
	conn.Close()
}

//func runMigration(conn *pgxpool.Pool, cfg *config.DBConfig) error {
//	if _, err := os.Stat(cfg.MigrationPath); os.IsNotExist(err) {
//		return fmt.Errorf("migration file not found: %s", cfg.MigrationPath)
//	}
//
//	content, err := os.ReadFile(cfg.MigrationPath)
//	if err != nil {
//		return fmt.Errorf("failed to read migration file: %w", err)
//	}
//
//	tx, err := conn.Begin(context.Background())
//	if err != nil {
//		return fmt.Errorf("failed to begin transaction: %w", err)
//	}
//	defer tx.Rollback(context.Background())
//
//	_, err = tx.Exec(context.Background(), string(content))
//	if err != nil {
//		return fmt.Errorf("failed to execute migration: %w", err)
//	}
//
//	err = tx.Commit(context.Background())
//	if err != nil {
//		return fmt.Errorf("failed to commit migration: %w", err)
//	}
//
//	fmt.Println("Migration completed successfully")
//	return nil
//}
