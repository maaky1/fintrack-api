package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			CREATE TABLE IF NOT EXISTS master.users (
				id BIGSERIAL PRIMARY KEY,
				clerk_user_id TEXT NOT NULL UNIQUE,
				fullname TEXT,
				created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
			);
		`)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			DROP TABLE IF EXISTS master.users;
		`)
		return err
	})
}
