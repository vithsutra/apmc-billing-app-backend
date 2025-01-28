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

type ConsigneeRepo struct {
	db *sql.DB
}

func NeeConsigneeRepo(db *sql.DB) *ConsigneeRepo {
	return &ConsigneeRepo{
		db: db,
	}
}

func (c *ConsigneeRepo) CreateConsignee(r *http.Request) error {
	var consignee models.Consignee

	if err := json.NewDecoder(r.Body).Decode(&consignee); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(consignee); err != nil {
		return err
	}

	consignee.ConsigneeId = uuid.NewString()

	query := database.NewQuery(c.db)
	if err := query.CreateConsignee(&consignee); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *ConsigneeRepo) DeleteConsignee(r *http.Request) error {
	vars := mux.Vars(r)

	consigneeId := vars["consignee_id"]

	if consigneeId == "" {
		return errors.New("Consignee id cannot empty")
	}

	query := database.NewQuery(c.db)
	if err := query.DeleteConsignee(consigneeId); err != nil {
		log.Println(err)
		return err
	}

	return nil

}
func (repo *ConsigneeRepo) GetConsignee(r *http.Request) ([]*models.Consignee, error) {
	vars := mux.Vars(r)

	userId := vars["user_id"]

	if userId == "" {
		return nil, errors.New("user id cannot be empty")
	}
	query := database.NewQuery(repo.db)

	consignees, err := query.GetConsignee(userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return consignees, nil
}
