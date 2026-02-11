package config

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *viper.Viper, zlog *zerolog.Logger) *gorm.DB {
	dsn := cfg.GetString("database.url")
	idleConnection := cfg.GetInt("database.pool.idle")
	maxConnection := cfg.GetInt("database.pool.max")
	maxLifeTimeConnection := cfg.GetInt("database.pool.lifetime_seconds")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed connect to database")
	}

	sqlDb, err := db.DB()
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to get sql.Db")
	}

	sqlDb.SetMaxIdleConns(idleConnection)
	sqlDb.SetMaxOpenConns(maxConnection)
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	zlog.Info().Msg("success connect to database")
	return db
}
