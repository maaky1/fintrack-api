package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
)

func RunMigration(dsn string, zlog *zerolog.Logger) {
	m, err := migrate.New(
		"file://internal/database/migrations",
		dsn,
	)

	if err != nil {
		zlog.Fatal().Err(err).Msg("migration init error")
	}

	err = m.Up()

	if err == migrate.ErrNoChange {
		zlog.Info().Msg("migration skipped (no changes)")
		return
	}

	if err != nil {
		zlog.Fatal().Err(err).Msg("migration run error")
	}

	zlog.Info().Msg("migration applied successfully")
}
