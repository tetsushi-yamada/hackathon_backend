package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/profile_picture"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ProfilePictureHandler struct {
	ProfilePictureUsecase *usecase.ProfilePictureUsecase
}

func NewProfilePictureHandler(profilePictureUsecase *usecase.ProfilePictureUsecase) *ProfilePictureHandler {
	return &ProfilePictureHandler{ProfilePictureUsecase: profilePictureUsecase}
}

func (h *ProfilePictureHandler) UploadProfilePictureHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("profile_picture")
	if err != nil {
		log.Printf("Error retrieving file: %v", err)
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			log.Printf("Error creating upload directory: %v", err)
			http.Error(w, "Unable to create upload directory", http.StatusInternalServerError)
			return
		}
	}

	// 一意のファイル名を生成
	uniqueFileName := uuid.New().String() + filepath.Ext(handler.Filename)
	destFile, err := os.Create(filepath.Join(uploadDir, uniqueFileName))
	if err != nil {
		log.Printf("Error creating destination file: %v", err)
		http.Error(w, "Unable to create destination file", http.StatusInternalServerError)
		return
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, file); err != nil {
		log.Printf("Error saving file: %v", err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	userID := r.FormValue("user_id")
	if userID == "" {
		log.Println("Invalid user ID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var profilePicture profile_picture.ProfilePicture
	profilePicture.UserID = userID
	profilePicture.ProfilePicture = destFile.Name()

	err = h.ProfilePictureUsecase.UploadProfilePicture(profilePicture)
	if err != nil {
		log.Printf("Error saving profile picture: %v", err)
		http.Error(w, "Unable to save profile picture", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully uploaded file: %v", handler.Filename)
}

func (h *ProfilePictureHandler) GetProfilePictureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	profile, err := h.ProfilePictureUsecase.GetProfilePicture(userID)
	if err != nil {
		if err.Error() == fmt.Sprintf("no profile picture found for user_id: %s", userID) {
			http.Error(w, "Profile picture not found", http.StatusNotFound)
			return
		}
		log.Printf("Error retrieving profile picture: %v", err)
		http.Error(w, "Unable to retrieve profile picture", http.StatusInternalServerError)
		return
	}

	file, err := os.Open(profile.ProfilePicture)
	if err != nil {
		http.Error(w, "Unable to open profile picture", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/jpeg") // 画像のMIMEタイプに応じて変更
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Unable to send profile picture", http.StatusInternalServerError)
	}
}
