package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/repository"
)

type BankerHandler struct {
	repo *repository.BankerRepo
}

func NewBankerHandler(repo *repository.BankerRepo) *BankerHandler {
	return &BankerHandler{repo: repo}
}

func (handler *BankerHandler) CreateBankerHandler(w http.ResponseWriter, r *http.Request) {
	if err := handler.repo.CreateBanker(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Banker created successfully"})

}

func (handler *BankerHandler) DeleteBankerHandler(w http.ResponseWriter, r *http.Request) {
	err := handler.repo.DeleteBanker(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Banker deleted successfully"})
}

func (handler *BankerHandler) GetBankerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bankers, err := handler.repo.GetBanker(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*models.Banker{"bankers": bankers})
}
