package tweet_picture

import "time"

type TweetPicture struct {
	TweetID      string    `json:"tweet_id" db:"tweet_id"`
	TweetPicture string    `json:"tweet_picture" db:"tweet_picture"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
