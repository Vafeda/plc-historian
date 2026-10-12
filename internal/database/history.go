package database

import (
	"context"

	"github.com/Vafeda/plc-historian/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type History struct {
	pool *pgxpool.Pool
}

func NewHistory(pool *pgxpool.Pool) *History {
	return &History{
		pool: pool,
	}
}

func (h *History) Write(ctx context.Context, raw, arch []models.Sample) error {
	tx, err := h.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.CopyFrom(ctx,
		pgx.Identifier{"history_raw"},
		[]string{"itemid", "itemtimestamp", "regtimestamp", "valint", "valdouble", "valbool", "valstring", "quality"},
		pgx.CopyFromSlice(len(raw), func(i int) ([]any, error) {
			s := raw[i]

			return []any{int32(s.ItemId), s.ItemTS, s.RegTS, s.ValInt, s.ValDouble, s.ValBool, s.ValStr, s.Quality}, nil
		}),
	)
	if err != nil {
		return err
	}

	_, err = tx.CopyFrom(ctx,
		pgx.Identifier{"history"},
		[]string{"itemid", "itemtimestamp", "regtimestamp", "valint", "valdouble", "valbool", "valstring", "quality"},
		pgx.CopyFromSlice(len(arch), func(i int) ([]any, error) {
			s := arch[i]

			return []any{int32(s.ItemId), s.ItemTS, s.RegTS, s.ValInt, s.ValDouble, s.ValBool, s.ValStr, s.Quality}, nil
		}),
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
