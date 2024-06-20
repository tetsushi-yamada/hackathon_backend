package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tetsushi-yamada/hackathon_backend/internal/domain/block"
	"github.com/tetsushi-yamada/hackathon_backend/internal/usecase"
	"log"
	"net/http"
)

type BlockHandler struct {
	BlockUsecase *usecase.BlockUsecase
}

func NewBlockHandler(bu *usecase.BlockUsecase) *BlockHandler {
	return &BlockHandler{BlockUsecase: bu}
}

func (bh *BlockHandler) CreateBlockHandler(w http.ResponseWriter, r *http.Request) {
	var block block.Block
	if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if block.UserID == block.BlockID {
		http.Error(w, "Cannot block yourself", http.StatusBadRequest)
		return
	}
	err := bh.BlockUsecase.CreateBlockUsecase(block)
	if err != nil {
		fmt.Printf("Failed to create block: %v\n", err)
		http.Error(w, "Failed to create block", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (bh *BlockHandler) GetBlockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	blockID := vars["block_id"]
	blockOrNot, err := bh.BlockUsecase.GetBlocksUsecase(userID, blockID)
	if err != nil {
		log.Printf("Failed to get block: %v\n", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(blockOrNot)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}

func (bh *BlockHandler) DeleteBlockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	blockID := vars["block_id"]
	err := bh.BlockUsecase.DeleteBlockUsecase(userID, blockID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusNoContent)
}
