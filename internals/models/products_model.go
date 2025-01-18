package models

type Products struct {
	ProductId   string `json:"product_id"`
	ProcuctName string `json:"product_name" validate:"required,max=50"`
	ProductHsn  string `json:"product_hsn" validate:"required,max=50"`
	ProductQty  string `json:"product_qty" validate:"required,max=50"`
	ProductUnit string `json:"product_unit" validate:"required,max=50"`
	ProductRate string `json:"product_rate" validate:"required,max=50"`
	InvoiceId   string `json:"invoice_id" validate:"required,max=100"`
}
