package config

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitLogger(cfg *viper.Viper) {
	zerolog.TimeFieldFormat = time.RFC3339Nano

	level := zerolog.InfoLevel

	app := cfg.GetString("app.name")
	env := strings.ToLower(cfg.GetString("app.env"))
	if env == "dev" || env == "local" {
		level = zerolog.DebugLevel

		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05",
		}

		log.Logger = zerolog.New(consoleWriter).
			Level(level).
			With().
			Timestamp().
			Str("app", app).
			Str("env", env).
			Logger()

		return
	}

	log.Logger = zerolog.New(os.Stdout).
		Level(level).
		With().
		Timestamp().
		Str("app", app).
		Str("env", env).
		Logger()
}
