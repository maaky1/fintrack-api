package models

import "time"

type UserModel struct {
	ID          int64     `bun:"id,pk,autoincrement"`
	ClerkUserID string    `bun:"clerk_user_id,notnull,unique"`
	Fullname    *string   `bun:"fullname,nullzero"`
	CreatedAt   time.Time `bun:"created_at,notnull"`
	UpdatedAt   time.Time `bun:"updated_at,notnull"`
}

func (UserModel) TableName() string {
	return "master.users"
}
