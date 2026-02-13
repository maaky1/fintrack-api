package service

import (
	"context"
	"fintrack-api/internal/app"
	"fintrack-api/internal/dto"
	"fintrack-api/internal/helper/logger"
	"fintrack-api/internal/models"
	"fintrack-api/internal/repository"
	"strings"
)

type UserService interface {
	CreateUser(ctx context.Context, u *dto.UserDto) (*dto.UserDtoResponse, error)
	GetByClerkUserID(ctx context.Context, clerkUserID string) (*dto.UserDtoResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, u *dto.UserDto) (*dto.UserDtoResponse, error) {
	zlog := logger.FromCtx(ctx, "service", "CreateUser")
	zlog.Info().Msg("in")

	if u == nil {
		return nil, app.InvalidInput("request body is required")
	}

	if strings.TrimSpace(u.ClerkUserID) == "" {
		return nil, app.InvalidInput("clerkUserId is required")
	}

	userModel := &models.UserModel{
		ClerkUserID: u.ClerkUserID,
		Fullname:    u.Fullname,
	}

	created, err := s.userRepo.InsertUser(ctx, userModel)
	if err != nil {
		zlog.Err(err).Msg("repository error")
		return nil, app.BadRequest("failed to create user")
	}

	zlog.Info().Msg("out")
	return &dto.UserDtoResponse{
		ID:          created.ID,
		ClerkUserID: created.ClerkUserID,
		Fullname:    u.Fullname,
		CreatedAt:   created.CreatedAt,
		UpdatedAt:   created.UpdatedAt,
	}, nil
}

func (s *userService) GetByClerkUserID(ctx context.Context, clerkUserID string) (*dto.UserDtoResponse, error) {
	zlog := logger.FromCtx(ctx, "service", "GetByClerkUserID")
	zlog.Info().Msg("in")

	user, err := s.userRepo.FindByClerkUserID(ctx, clerkUserID)
	if err != nil {
		zlog.Err(err).Msg("repository error")
		return nil, err
	}

	if user == nil {
		zlog.Warn().Msg("user not found")
		return nil, app.NotFound("user not found")
	}

	zlog.Info().Msg("out")
	return &dto.UserDtoResponse{
		ID:          user.ID,
		ClerkUserID: user.ClerkUserID,
		Fullname:    user.Fullname,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}
