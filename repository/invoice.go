package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
)

type InvoiceRepo struct {
	db *sql.DB
}

func NewInvoiceRepo(db *sql.DB) *InvoiceRepo {
	return &InvoiceRepo{
		db: db,
	}
}

func (repo *InvoiceRepo) CreateInvoice(r *http.Request) (string, error) {
	var invoice models.Invoice

	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		return "", errors.New("error occurred while decoding")
	}

	validate := validator.New()

	if err := validate.Struct(invoice); err != nil {
		return "", err
	}

	invoice.InvoiceId = uuid.NewString()
	invoice.PaymentStatus = false

	query := database.NewQuery(repo.db)

	if err := query.CreateInvoice(&invoice); err != nil {
		log.Println(err)
		return "", errors.New("error occurred with database")
	}

	return invoice.InvoiceId, nil
}

func (repo *InvoiceRepo) DeleteInvoice(r *http.Request) error {
	vars := mux.Vars(r)

	invoiceId := vars["invoice_id"]

	if invoiceId == "" {
		return errors.New("invoice id cannot be empty")
	}

	query := database.NewQuery(repo.db)

	if err := query.DeleteInvoice(invoiceId); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo *InvoiceRepo) GetInvoices(r *http.Request) ([]*models.InvoiceResponse, error) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	if userId == "" {
		return nil, errors.New("user id cannot be empty")
	}

	query := database.NewQuery(repo.db)

	invoices, err := query.GetInvoices(userId)

	if err != nil {
		log.Println(err)
		return nil, errors.New("error occurred with database")
	}

	return invoices, nil
}
