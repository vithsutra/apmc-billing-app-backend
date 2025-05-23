package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/utils"
)

type InvoiceHandler struct {
	invoiceRepo models.InvoiceInterface
}

func NewInvoiceHandler(invoiceRepo models.InvoiceInterface) *InvoiceHandler {
	return &InvoiceHandler{
		invoiceRepo: invoiceRepo,
	}
}

func (handler *InvoiceHandler) CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	invoiceId, err := handler.invoiceRepo.CreateInvoice(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"invoice_id": invoiceId})
}

func (handler *InvoiceHandler) DeleteInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := handler.invoiceRepo.DeleteInvoice(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "invoice deleted successfully"})
}

func (handler *InvoiceHandler) GetInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	invoices, err := handler.invoiceRepo.GetInvoices(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*models.InvoiceResponse{"invoices": invoices})
}

func (handler *InvoiceHandler) DownloadInvoiceHandler(w http.ResponseWriter, r *http.Request) {

	invoicePdf, err := handler.invoiceRepo.DownloadInvoice(r)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	if err := utils.GeneratePdf(w, invoicePdf); err != nil {
		log.Println("error occurred while generating the pdf, Error: ", err.Error())
	}
}

func (handler *InvoiceHandler) UpdateInvoicePaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	if err := handler.invoiceRepo.UpdateInvoiceStatus(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "payment status updated successfully"})
}
