package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	var User user.User
	newUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		http.Error(w, "Failed to generate UUID", http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	User.UserID = newUUID.String()
	err = uh.UserUsecase.CreateUserUsecase(User)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(User.UserID)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}

}

func (uh *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	User, err := uh.UserUsecase.GetUserUsecase(userID)
	if err != nil {
		http.Error(w, "User not found"+err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(User)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (uh *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	err := uh.UserUsecase.DeleteUserUsecase(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}

func (uh *UserHandler) SearchUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchWord := vars["search_word"]
	Users, err := uh.UserUsecase.SearchUserUsecase(searchWord)
	if err != nil {
		http.Error(w, "User not found"+err.Error(), http.StatusNotFound)
		return
	}
	Users_api := user.Users{Users: Users, Count: len(Users)}
	err = json.NewEncoder(w).Encode(Users_api)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}
