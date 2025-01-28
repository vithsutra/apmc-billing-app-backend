package models

import "net/http"

type Invoice struct {
	InvoiceId     string `json:"invoice_id"`
	Name          string `json:"name" validate:"required,max=100"`
	PaymentStatus bool   `json:"payment_status"`
	UserId        string `json:"user_id" validate:"required,max=100"`
	ReceiverId    string `json:"receiver_id" validate:"required,max=100"`
	ConsigneeId   string `json:"consignee_id" validate:"required,max=100"`
	InvoiceDate   string `json:"invoice_date" validate:"required,max=50"`
	SupplyDate    string `json:"supply_date" validate:"required,max=50"`
}

type InvoiceResponse struct {
	InvoiceId     string `json:"invoice_id" validate:"required,max=100"`
	Name          string `json:"name" validate:"required,max=100"`
	PaymentStatus bool   `json:"payment_status" validate:"required"`
}

type InvoiceInterface interface {
	CreateInvoice(*http.Request) (string, error)
	DeleteInvoice(*http.Request) error
	GetInvoices(*http.Request) ([]*InvoiceResponse, error)
}
