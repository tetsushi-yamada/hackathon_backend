package handler

type handlers struct {
	tweet    *tweetHandler
	user     *userHandler
	follow   *followHandler
	follower *followerHandler
}
