package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/user"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"net/http"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(uu *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: uu}
}

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user user.User
	newUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return
	}
	user.UserID = newUUID.String()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = uh.UserUsecase.CreateUserUsecase(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	user, err := uh.UserUsecase.GetUserUsecase(userID)
	if err != nil {
		http.Error(w, "User not found"+err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (uh *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	err := uh.UserUsecase.DeleteUserUsecase(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}
