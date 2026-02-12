package repository

import (
	"context"
	"fintrack-api/internal/models"
)

type UserRepository interface {
	FindByClerkUserID(ctx context.Context, clerkUserID string) (*models.UserModel, error)
}
