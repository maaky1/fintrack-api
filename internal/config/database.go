package config

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewDatabase(cfg *viper.Viper, zlog *zerolog.Logger) (*bun.DB, error) {
	dsn := cfg.GetString("database.url")
	idleConnection := cfg.GetInt("database.pool.idle")
	maxConnection := cfg.GetInt("database.pool.max")
	maxLifeTimeConnection := cfg.GetInt("database.pool.lifetime_seconds")
	debug := cfg.GetBool("database.debug")

	sqlDb, err := sql.Open("pgx", dsn)
	if err != nil {
		zlog.Error().Err(err).Msg("failed create database connection")
		return nil, err
	}

	sqlDb.SetMaxIdleConns(idleConnection)
	sqlDb.SetMaxOpenConns(maxConnection)
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	pingTimeout := 10 * time.Second
	pingCtx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	if err := sqlDb.PingContext(pingCtx); err != nil {
		_ = sqlDb.Close()
		return nil, err
	}

	db := bun.NewDB(sqlDb, pgdialect.New())

	if debug {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	}

	zlog.Info().
		Int("pool_idle", idleConnection).
		Int("pool_max", maxConnection).
		Int("pool_lifetime_seconds", maxLifeTimeConnection).
		Bool("debug_sql", debug).
		Msg("success connect to database")

	return db, nil
}
