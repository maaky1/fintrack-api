package dto

import "time"

type UserDto struct {
	ClerkUserID string  `json:"clerkUserId"`
	Fullname    *string `json:"fullname"`
}

type UserDtoResponse struct {
	ID          int64     `json:"id"`
	ClerkUserID string    `json:"clerkUserId"`
	Fullname    *string   `json:"fullname"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
