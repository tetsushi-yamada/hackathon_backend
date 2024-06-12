package follow

import "time"

type FollowRequest struct {
	UserID    string    `json:"user_id" db:"user_id"`
	FollowID  string    `json:"follow_id" db:"follow_id"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type FollowRequests struct {
	FollowRequests []*FollowRequest `json:"follow_requests"`
	Count          int              `json:"count"`
}

type FollowRequestOrNot struct {
	Bool bool `json:"bool"`
}
