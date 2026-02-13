package routes

import (
	"fintrack-api/internal/handler/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App         *fiber.App
	UserHandler *http.UserHandler
}

func (c *RouteConfig) Setup() {
	c.SetupRegister()
}

func (c *RouteConfig) SetupRegister() {
	api := c.App.Group("/api")

	user := api.Group("/user")
	user.Get("", c.UserHandler.GetByClerkUserID)
	user.Post("", c.UserHandler.CreateUser)

	api.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "Ok"})
	})
}
