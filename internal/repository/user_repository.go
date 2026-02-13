package repository

import (
	"context"
	"fintrack-api/internal/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, u *models.UserModel) (*models.UserModel, error)
	FindByClerkUserID(ctx context.Context, clerkUserID string) (*models.UserModel, error)
}
