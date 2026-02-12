package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			CREATE SCHEMA IF NOT EXISTS master;
		`)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		return nil
	})
}
