package handler

type Handlers struct {
	Tweet          *TweetHandler
	User           *UserHandler
	ProfilePicture *ProfilePictureHandler
	TweetPicture   *TweetPictureHandler
	Follow         *FollowHandler
	Follower       *FollowerHandler
	Good           *GoodHandler
}
