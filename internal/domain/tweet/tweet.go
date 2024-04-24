package domain

import (
	"time"
)

type Tweet struct {
	TweetId   string    `json:"tweet_id,omitempty"`
	UserId    string    `json:"user_id,omitempty"`
	TweetText string    `json:"tweet_text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
