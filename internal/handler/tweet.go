package handler

import (
	"encoding/json"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/tweet"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"net/http"
)

type TweetHandler struct {
	TweetUsecase *usecase.TweetUsecase
}

func NewTweetHandler(tu *usecase.TweetUsecase) *TweetHandler {
	return &TweetHandler{TweetUsecase: tu}
}

func (th *TweetHandler) CreateTweetHandler(w http.ResponseWriter, r *http.Request) {
	var tweet tweet.Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := th.TweetUsecase.CreateTweetUsecase(tweet)
	if err != nil {
		http.Error(w, "Failed to create tweet", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (th *TweetHandler) GetTweetsHandlerByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	tweets, err := th.TweetUsecase.GetTweetsUsecase(userID)
	if err != nil {
		http.Error(w, "Tweet not found", http.StatusNotFound)
		return
	}
	tweets_api := tweet.Tweets{Tweets: tweets, Count: len(tweets)}
	err = json.NewEncoder(w).Encode(tweets_api)
	if err != nil {
		http.Error(w, "Failed to encode tweet", http.StatusInternalServerError)
		return
	}
}

func (th *TweetHandler) DeleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("tweet_id")
	err := th.TweetUsecase.DeleteTweetUsecase(tweetID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}
