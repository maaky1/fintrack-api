package routes

import (
	"fintrack-api/internal/handler/http"
	"os"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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
	user.Post("", c.UserHandler.CreateUser)
	user.Get("", c.UserHandler.GetByClerkUserID)

	api.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "Ok"})
	})

	// Swagger hanya aktif di non-production
	if os.Getenv("ENV") != "prod" {

		c.App.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect("/swagger/index.html", fiber.StatusTemporaryRedirect)
		})

		c.App.Get("/swagger/*", fiberSwagger.WrapHandler)
	}
}
