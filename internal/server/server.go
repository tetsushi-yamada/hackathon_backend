package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/handler"
	"net/http"
	"strings"
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

	// /v1/users
	router.HandleFunc("/v1/users", handlers.User.CreateUserHandler).Methods("POST")
	router.HandleFunc("/v1/users", handlers.User.GetUserHandler).Methods("GET")
	router.HandleFunc("/v1/users", handlers.User.DeleteUserHandler).Methods("DELETE")

	// /v1/tweets
	router.HandleFunc("/v1/tweets", handlers.Tweet.CreateTweetHandler).Methods("POST")
	router.HandleFunc("/v1/tweets", handlers.Tweet.GetTweetsHandlerByUserID).Methods("GET")
	router.HandleFunc("/v1/tweets", handlers.Tweet.DeleteTweetHandler).Methods("DELETE")

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

	Route{
		"CreateFollow",
		strings.ToUpper("Post"),
		"/v1/follows",
		handler.CreateFollow,
	},
	Route{
		"GetFollowersForUser",
		strings.ToUpper("Get"),
		"/v1/followers",
		handler.GetFollowersForUser,
	},

	Route{
		"GetFollowsForUser",
		strings.ToUpper("Get"),
		"/v1/follows",
		handler.GetFollowsForUser,
	},
}
