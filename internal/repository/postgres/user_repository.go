package postgres

import (
	"context"
	"fintrack-api/internal/models"

	"github.com/uptrace/bun"
)

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByClerkUserID(ctx context.Context, clerkUserID string) (*models.UserModel, error)
