package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vsynclabs/billsoft/internals/models"
)

type ProductHandler struct {
	ProductRepo models.ProductInterface
}

func NewProductHandler(productRepo models.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductRepo: productRepo,
	}
}

func (handler *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := handler.ProductRepo.CreateProduct(r); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "error occurred"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "product created successfully"})
}

func (handler *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := handler.ProductRepo.DeleteProduct(r); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "error occurred"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "product deleted successfully"})
}

func (handler *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := handler.ProductRepo.GetProduct(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "error occurrred"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*models.Product{"receiver_details": products})
}
