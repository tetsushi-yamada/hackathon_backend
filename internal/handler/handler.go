package handler

type handlers struct {
	tweet    *TweetHandler
	user     *UserHandler
	follow   *FollowHandler
	follower *FollowerHandler
}
