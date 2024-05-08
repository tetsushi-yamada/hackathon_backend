package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/good"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"net/http"
)

type GoodHandler struct {
	GoodUsecase *usecase.GoodUsecase
}

func NewGoodHandler(gu *usecase.GoodUsecase) *GoodHandler {
	return &GoodHandler{GoodUsecase: gu}
}

func (gh *GoodHandler) CreateGoodHandler(w http.ResponseWriter, r *http.Request) {
	var good good.Good
	if err := json.NewDecoder(r.Body).Decode(&good); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := gh.GoodUsecase.CreateGoodUsecase(good)
	if err != nil {
		http.Error(w, "Failed to create good", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (gh *GoodHandler) GetGoodsHandler(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("tweet_id")
	userID := r.URL.Query().Get("user_id")

	if tweetID == "" && userID == "" {
		http.Error(w, "Both tweet_id and user_id cannot be empty", http.StatusBadRequest)
		return
	}

	var goods []*good.Good
	var err error

	if tweetID != "" {
		goods, err = gh.GoodUsecase.GetGoodsUsecaseByTweetID(tweetID)
	} else {
		goods, err = gh.GoodUsecase.GetGoodsUsecaseByUserID(userID)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	goodsAPI := good.Goods{Goods: goods, Count: len(goods)}
	if err = json.NewEncoder(w).Encode(goodsAPI); err != nil {
		http.Error(w, "Failed to encode good", http.StatusInternalServerError)
		return
	}
}

func (gh *GoodHandler) DeleteGoodHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["tweet_id"]
	userID := vars["user_id"]
	err := gh.GoodUsecase.DeleteGoodUsecase(tweetID, userID)
	if err != nil {
		http.Error(w, "Good not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
