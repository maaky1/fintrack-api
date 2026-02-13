package helper

import (
	"fintrack-api/internal/app"
	"fintrack-api/internal/response"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ParseUintParam(c *fiber.Ctx, key string) (uint, error) {
	raw := c.Params(key)
	n, err := strconv.Atoi(raw)
	if err != nil || n <= 0 {
		return 0, fiber.ErrBadRequest
	}

	return uint(n), nil
}

func ParseDateQuery(c *fiber.Ctx, key string) (string, error) {
	raw := c.Query(key)
	if raw == "" {
		return "", nil
	}

	if _, err := time.Parse("2006-01-02", raw); err != nil {
		return "", fiber.ErrBadRequest
	}

	return raw, nil
}

func WriteServiceError(ctx *fiber.Ctx, err error) error {
	appErr, ok := err.(*app.AppError)
	if !ok {
		return response.Error(ctx, fiber.StatusInternalServerError, "Internal error")
	}

	switch appErr.Code {
	case app.ErrBadRequest:
		return response.Error(ctx, fiber.StatusBadRequest, appErr.Message)
	case app.ErrInvalidInput:
		return response.Error(ctx, fiber.StatusBadRequest, appErr.Message)
	case app.ErrNotFound:
		return response.Error(ctx, fiber.StatusNotFound, appErr.Message)
	case app.ErrConflict:
		return response.Error(ctx, fiber.StatusConflict, appErr.Message)
	case app.ErrForbidden:
		return response.Error(ctx, fiber.StatusForbidden, appErr.Message)
	default:
		return response.Error(ctx, fiber.StatusInternalServerError, appErr.Message)
	}
}
