package models

import (
	"time"

	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:master.users"`

	ID          int64     `bun:"id,pk,autoincrement"`
	ClerkUserID string    `bun:"clerk_user_id,notnull,unique"`
	Fullname    *string   `bun:"fullname,nullzero"`
	CreatedAt   time.Time `bun:"created_at,nullzero,notnull,default:now()"`
	UpdatedAt   time.Time `bun:"updated_at,nullzero,notnull,default:now()"`
}
