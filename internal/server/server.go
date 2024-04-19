package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/handler"
	"github.com/tetsushi-yamada/hackathon_backend/internal/logger"
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

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

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
		"CreateTweet",
		strings.ToUpper("Post"),
		"/v1/tweets",
		handler.CreateTweet,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/v1/users",
		handler.CreateUser,
	},

	Route{
		"GetTweet",
		strings.ToUpper("Get"),
		"/v1/tweets/{tweet_id}",
		handler.GetTweet,
	},

	Route{
		"GetUser",
		strings.ToUpper("Get"),
		"/v1/users/{user_id}",
		handler.GetUser,
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
