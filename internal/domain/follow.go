package domain

type Follow struct {
	UserId   string `json:"user_id,omitempty"`
	FollowId string `json:"follow_id,omitempty"`
}
