package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/handler"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(handlers *handler.Handlers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// /v1/users,v1/user/{user_id}
	router.HandleFunc("/v1/users", handlers.User.CreateUserHandler).Methods("POST")
	router.HandleFunc("/v1/users/{user_id}", handlers.User.GetUserHandler).Methods("GET")
	router.HandleFunc("/v1/users/{user_id}", handlers.User.DeleteUserHandler).Methods("DELETE")

	// /v1/tweets
	router.HandleFunc("/v1/tweets", handlers.Tweet.CreateTweetHandler).Methods("POST")
	router.HandleFunc("/v1/tweets/{user_id}", handlers.Tweet.GetTweetsHandlerByUserID).Methods("GET")
	router.HandleFunc("/v1/tweets/{tweet_id}", handlers.Tweet.DeleteTweetHandler).Methods("DELETE")

	// /v1/follows
	router.HandleFunc("/v1/follows", handlers.Follow.CreateFollowHandler).Methods("POST")
	router.HandleFunc("/v1/follows/{user_id}", handlers.Follow.GetFollowsHandler).Methods("GET")
	router.HandleFunc("/v1/follows/{user_id}/{follow_id}", handlers.Follow.DeleteFollowHandler).Methods("DELETE")
	router.HandleFunc("/v1/follows/{user_id}/{follow_id}/check", handlers.Follow.GetFollowOrNotHandler).Methods("GET")

	// /v1/followers
	router.HandleFunc("/v1/followers/{follow_id}", handlers.Follower.GetFollowersHandler).Methods("GET")

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},
}
