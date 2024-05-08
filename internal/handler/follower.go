package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/follow"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"net/http"
)

type FollowerHandler struct {
	FollowerUsecase *usecase.FollowerUsecase
}

func NewFollowerHandler(fu *usecase.FollowerUsecase) *FollowerHandler {
	return &FollowerHandler{FollowerUsecase: fu}

}

func (fh *FollowerHandler) GetFollowersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	followID := vars["follow_id"]
	followers, err := fh.FollowerUsecase.GetFollowersUsecase(followID)
	if err != nil {
		fmt.Printf("Failed to get followers: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	followers_api := follow.Follows{Follows: followers, Count: len(followers)}
	err = json.NewEncoder(w).Encode(followers_api)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}
