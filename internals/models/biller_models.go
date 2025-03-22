package models

import "net/http"

type Biller struct {
	UserId            string `json:"user_id"`
	BillerId          string `json:"biller_id"`
	BillerName        string `json:"biller_name" validate:"required,max=50"`
	BillerAddress     string `json:"biller_address" validate:"required,max=100"`
	BillerMobile      string `json:"biller_mobile" validate:"required,max=50"`
	BillerGstin       string `json:"biller_gstin" validate:"required,max=50"`
	BillerPan         string `json:"biller_pan" validate:"required,max=50"`
	BillerMail        string `json:"biller_mail" validate:"required,max=50"`
	BillerCompanyLogo string `json:"biller_companylogo" `
}

type BillerInterface interface {
	CreateBiller(*http.Request) error
	DeleteBiller(*http.Request) error
	GetBiller(r *http.Request) ([]*Biller, error)
	UploadCompanyLogo(*http.Request) error
	DeleteCompanyLogo(*http.Request) error
}
