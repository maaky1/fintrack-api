package config

import (
	"fintrack-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func NewFiber(cfg *viper.Viper, zlog *zerolog.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: cfg.GetString("app.name"),
		Prefork: cfg.GetBool("app.prefork"),
	})

	app.Use(middleware.RequestLogger())

	return app
}
