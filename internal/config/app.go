package config

import (
	h "fintrack-api/internal/handler/http"
	"fintrack-api/internal/handler/routes"
	"fintrack-api/internal/repository/postgres"
	"fintrack-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
)

type BootstrapConfig struct {
	Config *viper.Viper
	Logger *zerolog.Logger
	DB     *bun.DB
	App    *fiber.App
}

func Bootstrap(cfg *BootstrapConfig) {
	userRepository := postgres.NewUserRepository(cfg.DB)
	userService := service.NewUserService(userRepository)
	userHandler := h.NewUserHandler(userService)

	routeConfig := routes.RouteConfig{
		App:         cfg.App,
		UserHandler: userHandler,
	}

	routeConfig.Setup()
}
