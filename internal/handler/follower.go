package handler

import "net/http"

type FollowerHandler struct{}

func GetFollowersForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Path: internal/handler/tweet.go
