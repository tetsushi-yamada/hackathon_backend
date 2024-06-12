package user

import (
	"time"
)

type User struct {
	UserID          string    `json:"user_id" db:"user_id"`
	UserName        string    `json:"user_name" db:"user_name"`
	UserDescription *string   `json:"user_description" db:"user_description"`
	IsPrivate       bool      `json:"is_private" db:"is_private"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type Users struct {
	Users []*User `json:"users"`
	Count int     `json:"count"`
}
