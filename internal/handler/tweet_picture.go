package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/tweet_picture"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type TweetPictureHandler struct {
	TweetPictureUsecase *usecase.TweetPictureUsecase
}

func NewTweetPictureHandler(TweetPictureUsecase *usecase.TweetPictureUsecase) *TweetPictureHandler {
	return &TweetPictureHandler{TweetPictureUsecase: TweetPictureUsecase}
}

func (tph *TweetPictureHandler) UploadTweetPictureHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("tweet_picture")
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

	tweetID := r.FormValue("tweet_id")
	if tweetID == "" {
		log.Println("Invalid tweet ID")
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}

	var tweetPicture tweet_picture.TweetPicture
	tweetPicture.TweetID = tweetID
	tweetPicture.TweetPicture = destFile.Name()

	err = tph.TweetPictureUsecase.UploadTweetPicture(tweetPicture)
	if err != nil {
		log.Printf("Error saving tweet picture: %v", err)
		http.Error(w, "Unable to save tweet picture", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully uploaded file: %v", handler.Filename)
}

func (tph *TweetPictureHandler) GetTweetPictureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["tweet_id"]

	tweet, err := tph.TweetPictureUsecase.GetTweetPicture(tweetID)
	if err != nil {
		if err.Error() == fmt.Sprintf("no profile picture found for tweet_id: %s", tweetID) {
			http.Error(w, "Tweet picture not found", http.StatusNotFound)
			return
		}
		log.Printf("Error retrieving tweet picture: %v", err)
		http.Error(w, "Unable to retrieve tweet picture", http.StatusInternalServerError)
		return
	}

	file, err := os.Open(tweet.TweetPicture)
	if err != nil {
		http.Error(w, "Unable to open tweet picture", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/jpeg") // 画像のMIMEタイプに応じて変更
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Unable to send tweet picture", http.StatusInternalServerError)
	}
}
