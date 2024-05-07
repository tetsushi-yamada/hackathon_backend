package good

import "time"

type Good struct {
	TweetID   string    `json:"tweet_id" db:"tweet_id"`
	UserID    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Goods struct {
	Goods []*Good `json:"goods"`
	Count int     `json:"count"`
}
