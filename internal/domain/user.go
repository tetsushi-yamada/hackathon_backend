package domain

import (
	"time"
)

type User struct {
	UserId    string    `json:"user_id,omitempty" db:"user_id"`
	UserName  string    `json:"user_name,omitempty" db:"user_name"`
	Email     string    `json:"email,omitempty" db:"email"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
