package models

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type InvoiceRequest struct {
	InvoiceName          string `json:"invoice_name" validate:"required,max=100"`
	InvoiceReverseCharge string `json:"invoice_reverse_charge" validate:"required,max=20"`
	InvoiceState         string `json:"invoice_state" validate:"required,max=50"`
	InvoiceStateCode     string `json:"invoice_state_code" validate:"required,max=20"`
	InvoiceChallanNumber string `json:"invoice_challan_number" validate:"required,max=50"`
	InvoiceVehicleNumber string `json:"invoice_vehicle_number" validate:"required,max=50"`
	InvoiceDateOfSupply  string `json:"invoice_date_of_supply" validate:"required,max=50"`
	InvoicePlaceOfSupply string `json:"invoice_place_of_supply" validate:"required,max=50"`
	InvoiceGst           string `json:"invoice_gst" validate:"required,max=50"`
	UserId               string `json:"user_id" validate:"required,max=100"`
	ReceiverId           string `json:"receiver_id" validate:"required,max=100"`
	ConsigneeId          string `json:"consignee_id" validate:"required,max=100"`
	BillerId             string `json:"biller_id" validate:"required,max=100"`
}

type InvoiceResponse struct {
	InvoiceId     string `json:"invoice_id" validate:"required,max=100"`
	Name          string `json:"name" validate:"required,max=100"`
	PaymentStatus bool   `json:"payment_status" validate:"required"`
}

type Invoice struct {
	InvoiceId              string
	InvoiceName            string
	InvoicePaymentStatus   bool
	InvoiceReverseRecharge string
	InvoiceNumber          int32
	InvoiceDate            string
	InvoiceState           string
	InvoiceStateCode       string
	InvoiceChallanNumber   string
	InvoiceVehicleNumber   string
	InvoiceDateOfSupply    string
	InvoicePlaceOfSupply   string
	InvoiceGst             string
	UserId                 string
	BilledId               string
	ShippedId              string
	BillerId               string
}

func NewInvoice(request *InvoiceRequest) *Invoice {
	invoiceId := uuid.New().String()
	invoicePaymentStatus := false
	invoiceDate := time.Now().Format("02/01/2006")

	return &Invoice{
		InvoiceId:              invoiceId,
		InvoiceName:            request.InvoiceName,
		InvoicePaymentStatus:   invoicePaymentStatus,
		InvoiceReverseRecharge: request.InvoiceReverseCharge,
		InvoiceDate:            invoiceDate,
		InvoiceState:           request.InvoiceState,
		InvoiceStateCode:       request.InvoiceStateCode,
		InvoiceChallanNumber:   request.InvoiceChallanNumber,
		InvoiceVehicleNumber:   request.InvoiceVehicleNumber,
		InvoiceDateOfSupply:    request.InvoiceDateOfSupply,
		InvoicePlaceOfSupply:   request.InvoicePlaceOfSupply,
		InvoiceGst:             request.InvoiceGst,
		UserId:                 request.UserId,
		BilledId:               request.ReceiverId,
		ShippedId:              request.ConsigneeId,
		BillerId:               request.BillerId,
	}
}

type InvoicePdf struct {
	UserName             string
	UserAddress          string
	UserPhone            string
	UserEmail            string
	UserGstin            string
	UserPan              string
	InvoiceReverseCharge string
	InvoiceNumber        string
	InvoiceDate          string
	InvoiceState         string
	InvoiceStateCode     string
	InvoiceChallanNumber string
	InvoiceVehicleNumber string
	InvoiceDateOfSupply  string
	InvoicePlaceOfSupply string
	InvoiceGst           string
	ReceiverName         string
	ReceiverAdddress     string
	ReceiverGstin        string
	ReceiverState        string
	ReceiverStateCode    string
	ConsigneeName        string
	ConsigneeAddress     string
	ConsigneeGstin       string
	ConsigneeMobile      string
	ConsigneeState       string
	ConsigneeStateCode   string
	TotalQty             string
	GrandTotal           string
	Products             []*ProductPdf
}

type InvoiceInterface interface {
	CreateInvoice(r *http.Request) (string, error)
	DeleteInvoice(r *http.Request) error
	GetInvoices(r *http.Request) ([]*InvoiceResponse, error)
	UpdateInvoiceStatus(r *http.Request) error
	DownloadInvoice(r *http.Request) (*InvoicePdf, error)
}
