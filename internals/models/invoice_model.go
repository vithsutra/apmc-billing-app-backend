package models

type Invoice struct {
	InvoiceId   string `json:"invoice_id"`
	UserId      string `json:"user_id" validate:"required,max=100"`
	ReceiverId  string `json:"receiver_id" validate:"required,max=100"`
	ConsigneeId string `json:"consignee_id" validate:"required,max=100"`
	InvoiceDate string `json:"invoice_date" validate:"required,max=50"`
	SupplyDate  string `json:"supply_date" validate:"required,max=50"`
}
