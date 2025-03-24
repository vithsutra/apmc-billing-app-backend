package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vsynclabs/billsoft/internals/models"
)

type BillerHandler struct {
	BillerRepo models.BillerInterface
}

func NewBillerHandler(BillerRepo models.BillerInterface) *BillerHandler {
	return &BillerHandler{
		BillerRepo: BillerRepo,
	}
}

func (bh *BillerHandler) CreateBillerHandler(w http.ResponseWriter, r *http.Request) {
	err := bh.BillerRepo.CreateBiller(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Biller created successfully"})
}

func (bh *BillerHandler) DeleteBillerHandler(w http.ResponseWriter, r *http.Request) {
	err := bh.BillerRepo.DeleteBiller(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Biller deleted successfully"})
}

func (bh *BillerHandler) GetBillerHandler(w http.ResponseWriter, r *http.Request) {
	res, err := bh.BillerRepo.GetBiller(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (bh *BillerHandler) UploadCompanyLogoHandler(w http.ResponseWriter, r *http.Request) {
	err := bh.BillerRepo.UploadCompanyLogo(r)
	if err != nil {
		log.Println("Upload error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Company logo uploaded successfully"})
}

func (bh *BillerHandler) DeleteCompanyLogoHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("fileName")
	if fileName == "" {
		http.Error(w, "fileName is required", http.StatusBadRequest)
		return
	}
	err := bh.BillerRepo.DeleteCompanyLogo(r)
	if err != nil {
		log.Println("Delete error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Company logo deleted successfully"})
}
