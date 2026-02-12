package main

import (
	"fintrack-api/internal/config"
	"fintrack-api/internal/database"
	"fmt"
	"net/http"

	zlog "github.com/rs/zerolog/log"
)

func main() {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to load config")
	}

	// init logger (GLOBAL)
	config.InitLogger(cfg)

	// init database
	db, err := config.NewDatabase(cfg, &zlog.Logger)
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed connect to database")
	}
	defer db.Close()

	// run migration
	if err := database.RunMigration(db, &zlog.Logger); err != nil {
		zlog.Fatal().Err(err).Msg("failed migration")
	}

	// ambil port dari config
	port := cfg.GetString("app.port")

	// log startup
	zlog.Info().Str("port", port).Msg("server starting")

	// start server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		zlog.Fatal().Err(err).Msg("failed to start server")
	}
}
