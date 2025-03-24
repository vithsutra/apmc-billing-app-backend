package models

import "net/http"

type Banker struct {
	UserId            string `json:"user_id"`
	BankId            string `json:"bank_id"`
	BankName          string `json:"bank_name" validate:"required,max=50"`
	BankBranch        string `json:"bank_branch" validate:"required,max=50"`
	BankAccountNumber string `json:"bank_account_number" validate:"required,max=50"`
	BankIfscCode      string `json:"bank_ifsc_code" validate:"required,max=50"`
	BankHolderName    string `json:"bank_holder_name" validate:"required,max=50"`
}

type BankerInterface interface {
	CreateBanker(*http.Request) error
	DeleteBanker(*http.Request) error
	GetBanker(*http.Request) ([]*Banker, error)
}
