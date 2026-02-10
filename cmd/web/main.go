package main

import (
	"fintrack-api/internal/config"
	"fmt"
	"net/http"

	zlog "github.com/rs/zerolog/log"
)

func main() {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		zlog.Fatal().
			Err(err).
			Msg("failed to load config")
	}

	// init logger (GLOBAL)
	config.Init(cfg)

	// ambil port dari config
	port := cfg.GetString("app.port")
	if port == "" {
		port = "8080"
	}

	// log startup
	zlog.Info().
		Str("port", port).
		Msg("server starting")

	// 6. start server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		zlog.Fatal().
			Err(err).
			Msg("failed to start server")
	}
}
