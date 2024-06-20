package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	newUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return
	}
	tweet.TweetID = newUUID.String()
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = th.TweetUsecase.CreateTweetUsecase(tweet)
	if err != nil {
		http.Error(w, "Failed to create tweet", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(tweet.TweetID)
	if err != nil {
		http.Error(w, "Failed to encode tweet", http.StatusInternalServerError)
		return
	}
}

func (th *TweetHandler) GetTweetsHandlerByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
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

func (th *TweetHandler) GetTweetByTweetIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["tweet_id"]
	tweet, err := th.TweetUsecase.GetTweetByTweetIDUsecase(tweetID)
	if err != nil {
		http.Error(w, "Tweet not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(tweet)
	if err != nil {
		http.Error(w, "Failed to encode tweet", http.StatusInternalServerError)
		return
	}
}

func (th *TweetHandler) UpdateTweetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var tweet tweet.Tweet
	tweetID := vars["tweet_id"]
	tweet.TweetID = tweetID
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := th.TweetUsecase.UpdateTweetUsecase(tweet)
	if err != nil {
		http.Error(w, "Failed to update tweet", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (th *TweetHandler) DeleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["tweet_id"]
	err := th.TweetUsecase.DeleteTweetUsecase(tweetID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (th *TweetHandler) SearchTweetsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchWord := vars["search_word"]
	tweets, err := th.TweetUsecase.SearchTweetsUsecase(searchWord)
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

func (th *TweetHandler) GetRepliesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["tweet_id"]
	replies, err := th.TweetUsecase.GetRepliesUsecase(tweetID)
	if err != nil {
		http.Error(w, "Tweet not found", http.StatusNotFound)
		return
	}
	replies_api := tweet.Tweets{Tweets: replies, Count: len(replies)}
	err = json.NewEncoder(w).Encode(replies_api)
	if err != nil {
		http.Error(w, "Failed to encode tweet", http.StatusInternalServerError)
		return
	}
}

func (th *TweetHandler) GetRepostsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["tweet_id"]
	reposts, err := th.TweetUsecase.GetRepostsUsecase(tweetID)
	if err != nil {
		http.Error(w, "Tweet not found", http.StatusNotFound)
		return
	}
	reposts_api := tweet.Tweets{Tweets: reposts, Count: len(reposts)}
	err = json.NewEncoder(w).Encode(reposts_api)
	if err != nil {
		http.Error(w, "Failed to encode tweet", http.StatusInternalServerError)
		return
	}
}
