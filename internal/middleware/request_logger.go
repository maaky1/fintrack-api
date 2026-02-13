package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}

		start := time.Now()

		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Set("X-Request-ID", requestID)

		// child logger dari global logger (app/env kebawa)
		l := log.Logger.With().
			Str("transport", "http").
			Str("request_id", requestID).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Logger()

		// inject ke UserContext (Fiber-native)
		ctx := l.WithContext(c.UserContext())
		c.SetUserContext(ctx)

		// IN
		l.Info().Msg("request_started")

		err := c.Next()
		status := c.Response().StatusCode()
		took := time.Since(start)

		event := l.Info()
		if status >= 500 {
			event = l.Error()
		} else if status >= 400 {
			event = l.Warn()
		}

		// OUT
		event.
			Int("status", status).
			Dur("took", took).
			Err(err).
			Msg("request_completed")

		return err
	}
}
