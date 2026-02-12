package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS master.categories (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			name TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
			CONSTRAINT fk_categories_user
				FOREIGN KEY (user_id)
				REFERENCES master.users(id)
				ON DELETE CASCADE
		);

		CREATE UNIQUE INDEX IF NOT EXISTS uq_categories_user_id_name
			ON master.categories(user_id, name);
		`)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			DROP TABLE IF EXISTS master.categories;
		`)
		return err
	})
}
