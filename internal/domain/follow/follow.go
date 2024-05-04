package follow

import "time"

type Follow struct {
	UserID    string    `json:"user_id"`
	FollowID  string    `json:"follow_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Follows struct {
	Follows []*Follow `json:"follows"`
	Count   int       `json:"count"`
}

type FollowOrNot struct {
	Bool bool `json:"bool"`
}
