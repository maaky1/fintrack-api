package main

import (
	"fintrack-api/internal/config"
	"fintrack-api/internal/database"
	"fmt"

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

	// Init Fiber
	app := config.NewFiber(cfg, &zlog.Logger)

	// Register routes / wiring
	config.Bootstrap(&config.BootstrapConfig{
		Config: cfg,
		Logger: &zlog.Logger,
		DB:     db,
		App:    app,
	})

	// ambil port dari config
	port := cfg.GetString("app.port")

	// log startup
	zlog.Info().Str("port", port).Msg("server starting")

	// start server
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		zlog.Fatal().Err(err).Msg("failed to start server")
	}
}
