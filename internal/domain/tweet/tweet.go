package tweet

import (
	"time"
)

type Tweet struct {
	TweetID         string    `json:"tweet_id" db:"tweet_id"`
	UserID          string    `json:"user_id" db:"user_id"`
	TweetText       string    `json:"tweet_text" db:"tweet_text"`
	ParentID        *string   `json:"parent_id" db:"parent_id"`
	RetweetID       *string   `json:"retweet_id" db:"retweet_id"`
	IsInappropriate bool      `json:"is_inappropriate" db:"is_inappropriate"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type Tweets struct {
	Tweets []*Tweet `json:"tweets"`
	Count  int      `json:"count"`
}
