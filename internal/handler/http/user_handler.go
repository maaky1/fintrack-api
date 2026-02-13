package http

import (
	"fintrack-api/internal/dto"
	"fintrack-api/internal/helper"
	"fintrack-api/internal/helper/logger"
	"fintrack-api/internal/response"
	"fintrack-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userSvc service.UserService
}

func NewUserHandler(userSvc service.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	ctx := c.UserContext()

	zlog := logger.FromCtx(ctx, "handler", "CreateUser")
	zlog.Info().Msg("in")

	var req dto.UserDto
	if err := c.BodyParser(&req); err != nil {
		zlog.Err(err).Msg("invalid body request")
		return response.Error(c, fiber.StatusBadRequest, "invalid body request")
	}

	created, err := h.userSvc.CreateUser(ctx, &req)
	if err != nil {
		zlog.Err(err).Msg("service error")
		return helper.WriteServiceError(c, err)
	}

	zlog.Info().Msg("out")
	return response.Success(c, fiber.StatusCreated, "user created", created)
}

func (h *UserHandler) GetByClerkUserID(c *fiber.Ctx) error {
	ctx := c.UserContext()

	zlog := logger.FromCtx(ctx, "handler", "GetByClerkUserID")
	zlog.Info().Msg("in")

	clerkUserID := c.Get("X-Clerk-User-Id")
	if clerkUserID == "" {
		zlog.Info().Msg("missing clerk user id")
		return response.Error(c, fiber.StatusBadRequest, "X-Clerk-User-Id header required")
	}

	user, err := h.userSvc.GetByClerkUserID(ctx, clerkUserID)
	if err != nil {
		zlog.Err(err).Msg("service error")
		return helper.WriteServiceError(c, err)
	}

	zlog.Info().Msg("out")
	return response.Success(c, fiber.StatusOK, "user found", user)
}
