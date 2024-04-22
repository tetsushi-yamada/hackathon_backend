package handler

type Handlers struct {
	Tweet    *TweetHandler
	User     *UserHandler
	Follow   *FollowHandler
	Follower *FollowerHandler
}
