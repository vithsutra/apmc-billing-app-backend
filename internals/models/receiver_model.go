package models

import "net/http"

type Receiver struct {
	ReceiverId        string `json:"receiver_id"`
	ReceiverName      string `json:"receiver_name" validate:"required,max=50"`
	ReceiverAddress   string `json:"receiver_address" validate:"required,max=100"`
	ReceiverGstin     string `json:"receiver_gstin" validate:"required,max=50"`
	ReceiverState     string `json:"receiver_state" validate:"required,max=50"`
	ReceiverStateCode string `json:"receiver_state_code" validate:"required,max=50"`
	UserId            string `json:"user_id" validate:"required,max=100"`
}

type ReceiverInterface interface {
	CreateReceiver(*http.Request) error
	DeleteReceiver(*http.Request) error
	GetReceivers(*http.Request) ([]*Receiver, error)
}
