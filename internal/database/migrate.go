package database

import (
	"context"
	"fintrack-api/internal/database/migrations"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func RunMigration(db *bun.DB, zlog *zerolog.Logger) error {
	ctx := context.Background()

	migrator := migrate.NewMigrator(db, migrations.Migrations)
	if err := migrator.Init(ctx); err != nil {
		return err
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	var groupMigrations = len(group.Migrations)
	if groupMigrations == 0 {
		zlog.Info().Int("migrations_applied", groupMigrations).Msg("no new migrations to apply")
		return nil
	}

	zlog.Info().
		Int64("migration_group", group.ID).
		Int("migrations_applied", groupMigrations).
		Msg("migration applied successfully")
	return nil
}
