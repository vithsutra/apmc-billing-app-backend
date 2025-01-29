package models

import "net/http"

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
