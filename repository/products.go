package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
	"gopkg.in/validator.v2"
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
	var product models.Products

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return err
	}

	if err := validator.Validate(product); err != nil {
		return err
	}
	product.ProductId = uuid.NewString()
	query := database.NewQuery(p.db)
	if err := query.CreateProduct(&product); err != nil {
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
