package profile_picture

import "time"

type ProfilePicture struct {
	UserID         string    `json:"user_id" db:"user_id"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
