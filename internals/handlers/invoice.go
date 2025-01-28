package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vsynclabs/billsoft/internals/models"
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"invoice_id": invoiceId})
}

func (handler *InvoiceHandler) DeleteInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	if err := handler.invoiceRepo.DeleteInvoice(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user deleted successfully"})
}

func (handler *InvoiceHandler) GetInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	invoices, err := handler.invoiceRepo.GetInvoices(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*models.InvoiceResponse{"invoices": invoices})
}
