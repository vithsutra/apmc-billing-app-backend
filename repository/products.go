package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) CreateProduct(r *http.Request) error {
	var productRequest models.ProductRequest

	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(productRequest); err != nil {
		return err
	}

	product, err := models.NewProduct(&productRequest)

	if err != nil {
		log.Println(err)
		return errors.New("error converting string to number")
	}

	query := database.NewQuery(p.db)
	if err := query.CreateProduct(product); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *ProductRepo) DeleteProduct(r *http.Request) error {
	vars := mux.Vars(r)
	productId := vars["product_id"]

	if productId == "" {
		return errors.New("product id Can't be empty")
	}
	query := database.NewQuery(p.db)

	if err := query.DeleteProduct(productId); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *ProductRepo) GetProduct(r *http.Request) ([]*models.Product, error) {
	vars := mux.Vars(r)
	invoiceId := vars["invoice_id"]

	if invoiceId == "" {
		return nil, errors.New("invoice  id  cannot be empty")
	}
	query := database.NewQuery(p.db)

	products, err := query.GetProduct(invoiceId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
}
