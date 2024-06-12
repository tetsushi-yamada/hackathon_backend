package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"log"
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
	vars := mux.Vars(r)
	userID := vars["user_id"]
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
	vars := mux.Vars(r)
	userID := vars["user_id"]
	followID := vars["follow_id"]
	err := fh.FollowUsecase.DeleteFollowUsecase(userID, followID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}

func (fh *FollowHandler) GetFollowOrNotHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	followID := vars["follow_id"]
	follow_bool, err := fh.FollowUsecase.GetFollowOrNotUsecase(userID, followID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(follow_bool)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (fh *FollowHandler) UpdateFollowRequestHandler(w http.ResponseWriter, r *http.Request) {
	var followRequest follow.FollowRequest
	if err := json.NewDecoder(r.Body).Decode(&followRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := fh.FollowUsecase.UpdateFollowRequestUsecase(followRequest)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (fh *FollowHandler) GetFollowRequestsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["follow_id"]
	followRequests, err := fh.FollowUsecase.GetFollowRequestsUsecase(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	followRequestsApi := follow.FollowRequests{FollowRequests: followRequests, Count: len(followRequests)}
	err = json.NewEncoder(w).Encode(followRequestsApi)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (fh *FollowHandler) GetFollowRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	followID := vars["follow_id"]
	followRequest, err := fh.FollowUsecase.GetFollowRequestUsecase(userID, followID)
	if err != nil {
		log.Print(err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(followRequest)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}
