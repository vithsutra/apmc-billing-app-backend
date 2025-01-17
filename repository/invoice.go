package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
	"gopkg.in/validator.v2"
)

type InvoiceRepo struct {
	db *sql.DB
}

func NewInvoiceRepo(db *sql.DB) *InvoiceRepo {
	return &InvoiceRepo{
		db: db,
	}
}

func (repo *InvoiceRepo) CreateInvoice(r *http.Request) error {
	var invoice models.Invoice

	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		return err
	}

	if err := validator.Validate(invoice); err != nil {
		return err
	}

	invoice.InvoiceId = uuid.NewString()

	query := database.NewQuery(repo.db)

	if err := query.CreateInvoice(&invoice); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
