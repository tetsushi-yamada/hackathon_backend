package handler

import (
	"encoding/json"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"net/http"
)

type FollowHandler struct {
	FollowUsecase *usecase.FollowUsecase
}

func NewFollowHandler(fu *usecase.FollowUsecase) *FollowHandler {
	return &FollowHandler{FollowUsecase: fu}
}

func (fh *FollowHandler) CreateFollowHandler(w http.ResponseWriter, r *http.Request) {
	var follow follow.Follow
	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if follow.UserID == follow.FollowID {
		http.Error(w, "Cannot follow yourself", http.StatusBadRequest)
		return
	}
	err := fh.FollowUsecase.CreateFollowUsecase(follow)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (fh *FollowHandler) GetFollowsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	follows, err := fh.FollowUsecase.GetFollowsUsecase(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	follows_api := follow.Follows{Follows: follows, Count: len(follows)}
	err = json.NewEncoder(w).Encode(follows_api)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (fh *FollowHandler) DeleteFollowHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	followID := r.URL.Query().Get("follow_id")
	err := fh.FollowUsecase.DeleteFollowUsecase(userID, followID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}
