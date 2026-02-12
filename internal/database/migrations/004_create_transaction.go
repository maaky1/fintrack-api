package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				CREATE TABLE IF NOT EXISTS master.transactions (
					id BIGSERIAL PRIMARY KEY,
					user_id BIGINT NOT NULL,
					category_id BIGINT NOT NULL,
					type TEXT NOT NULL CHECK (type IN ('income','expense')),
					amount NUMERIC(14,2) NOT NULL,
					description TEXT,
					occurred_at TIMESTAMPTZ NOT NULL,
					created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
					updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
					CONSTRAINT fk_transactions_user
						FOREIGN KEY (user_id)
						REFERENCES master.users(id)
						ON DELETE CASCADE,
					CONSTRAINT fk_transactions_category
						FOREIGN KEY (category_id)
						REFERENCES master.categories(id)
						ON DELETE RESTRICT
				);
			`)
			return err
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				DROP TABLE IF EXISTS master.transactions;
			`)
			return err
		},
	)
}
