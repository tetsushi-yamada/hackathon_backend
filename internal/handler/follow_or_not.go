package handler

import (
	"encoding/json"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"net/http"
)

type FollowOrNotHandler struct {
	FollowOrNotUsecase *usecase.FollowOrNotUsecase
}

func NewFollowOrNotHandler(fu *usecase.FollowOrNotUsecase) *FollowOrNotHandler {
	return &FollowOrNotHandler{FollowOrNotUsecase: fu}
}

func (fh *FollowOrNotHandler) GetFollowOrNotHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	followID := r.URL.Query().Get("follow_id")
	follow_bool, err := fh.FollowOrNotUsecase.GetFollowOrNotUsecase(userID, followID)
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
