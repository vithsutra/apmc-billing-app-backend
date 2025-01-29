package models

import (
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type ProductRequest struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name" validate:"required,max=50"`
	ProductHsn  string `json:"product_hsn" validate:"required,max=50"`
	ProductQty  string `json:"product_qty" validate:"required,max=50"`
	ProductUnit string `json:"product_unit" validate:"required,max=50"`
	ProductRate string `json:"product_rate" validate:"required,max=50"`
	InvoiceId   string `json:"invoice_id" validate:"required,max=100"`
}

type Product struct {
	ProductId   string
	ProductName string
	ProductHsn  string
	ProductQty  string
	ProductUnit string
	ProductRate string
	InvoiceId   string
	Total       string
}

func NewProduct(request *ProductRequest) (*Product, error) {
	productId := uuid.New().String()

	productQty, err := strconv.Atoi(request.ProductQty)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	productRate, err := strconv.Atoi(request.ProductRate)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	productTotal := productQty * productRate

	return &Product{
		ProductId:   productId,
		ProductName: request.ProductName,
		ProductHsn:  request.ProductHsn,
		ProductQty:  request.ProductQty,
		ProductUnit: request.ProductUnit,
		ProductRate: request.ProductRate,
		InvoiceId:   request.InvoiceId,
		Total:       strconv.Itoa(productTotal),
	}, nil

}

type ProductPdf struct {
	ProductName string
	ProductHsn  string
	ProductQty  string
	ProductUnit string
	ProductRate string
	Total       string
}

type ProductInterface interface {
	CreateProduct(*http.Request) error
	DeleteProduct(*http.Request) error
}
