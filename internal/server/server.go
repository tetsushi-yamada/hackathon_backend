package server

import (
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/handler"
)

func NewRouter(handlers *handler.Handlers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// /v1/users,v1/user/{user_id}
	router.HandleFunc("/v1/users", handlers.User.CreateUserHandler).Methods("POST")
	router.HandleFunc("/v1/users/{user_id}", handlers.User.GetUserHandler).Methods("GET")
	router.HandleFunc("/v1/users/{user_id}", handlers.User.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/v1/users/search/{search_word}", handlers.User.SearchUsersHandler).Methods("GET")

	// /v1/tweets
	router.HandleFunc("/v1/tweets", handlers.Tweet.CreateTweetHandler).Methods("POST")
	router.HandleFunc("/v1/tweets/{user_id}", handlers.Tweet.GetTweetsHandlerByUserID).Methods("GET")
	router.HandleFunc("/v1/tweets/by-tweet/{tweet_id}", handlers.Tweet.DeleteTweetHandler).Methods("DELETE")

	// /v1/follows
	router.HandleFunc("/v1/follows", handlers.Follow.CreateFollowHandler).Methods("POST")
	router.HandleFunc("/v1/follows/{user_id}", handlers.Follow.GetFollowsHandler).Methods("GET")
	router.HandleFunc("/v1/follows/{user_id}/{follow_id}", handlers.Follow.DeleteFollowHandler).Methods("DELETE")
	router.HandleFunc("/v1/follows/{user_id}/{follow_id}/check", handlers.Follow.GetFollowOrNotHandler).Methods("GET")

	// /v1/followers
	router.HandleFunc("/v1/followers/{follow_id}", handlers.Follower.GetFollowersHandler).Methods("GET")

	// /v1/goods
	router.HandleFunc("/v1/goods", handlers.Good.CreateGoodHandler).Methods("POST")
	router.HandleFunc("/v1/goods", handlers.Good.GetGoodsHandler).Methods("GET")
	router.HandleFunc("/v1/goods/{tweet_id}/{user_id}", handlers.Good.DeleteGoodHandler).Methods("DELETE")

	return router
}
