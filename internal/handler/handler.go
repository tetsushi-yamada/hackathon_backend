package handler

type Handlers struct {
	Tweet          *TweetHandler
	User           *UserHandler
	ProfilePicture *ProfilePictureHandler
	Follow         *FollowHandler
	Follower       *FollowerHandler
	Good           *GoodHandler
}
