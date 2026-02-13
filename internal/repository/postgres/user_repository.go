package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fintrack-api/internal/helper/logger"
	"fintrack-api/internal/models"

	"github.com/uptrace/bun"
)

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) InsertUser(ctx context.Context, u *models.UserModel) (*models.UserModel, error) {
	zlog := logger.FromCtx(ctx, "repository", "InsertUser")
	zlog.Info().Msg("in")

	exist := new(models.UserModel)
	err := r.db.NewSelect().
		Model(exist).
		Where("clerk_user_id = ?", u.ClerkUserID).
		Limit(1).
		Scan(ctx)

	if err == nil {
		zlog.Info().Str("result", "user already exist").Msg("conflict")
		return exist, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		zlog.Err(err).Msg("select failed")
		return nil, err
	}

	_, err = r.db.NewInsert().
		Model(u).
		Returning("*").
		Exec(ctx)

	if err != nil {
		zlog.Err(err).Msg("insert failed")
		return nil, err
	}

	zlog.Info().Msg("out")
	return u, nil
}

func (r *userRepository) FindByClerkUserID(ctx context.Context, clerkUserID string) (*models.UserModel, error) {
	zlog := logger.FromCtx(ctx, "repository", "FindByClerkUserID")
	zlog.Info().Msg("in")

	user := new(models.UserModel)
	err := r.db.NewSelect().
		Model(user).
		Where("clerk_user_id = ?", clerkUserID).
		Limit(1).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zlog.Err(err).Msg("not found")
			return nil, nil
		}
		zlog.Err(err).Msg("database error")
		return nil, err
	}

	zlog.Info().Msg("out")
	return user, nil
}
