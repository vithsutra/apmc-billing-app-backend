package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
)

type BankerRepo struct {
	db *sql.DB
}

func NewBankerRepo(db *sql.DB) *BankerRepo {
	return &BankerRepo{db: db}
}

func (b *BankerRepo) CreateBanker(r *http.Request) error {
	var banker models.Banker

	if err := json.NewDecoder(r.Body).Decode(&banker); err != nil {
		return fmt.Errorf("invalid request payload: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(banker); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	banker.BankId = uuid.NewString()

	query := database.NewQuery(b.db)
	if err := query.CreateBanker(&banker); err != nil {
		log.Println("Database error:", err)
		return fmt.Errorf("failed to create banker: %w", err)
	}
	return nil
}

func (b *BankerRepo) DeleteBanker(r *http.Request) error {
	vars := mux.Vars(r)
	bankerId, ok := vars["banker_id"]
	if !ok || bankerId == "" {
		return errors.New("banker_id cannot be empty")
	}

	query := database.NewQuery(b.db)
	if err := query.DeleteBanker(bankerId); err != nil {
		log.Println("Database error:", err)
		return fmt.Errorf("failed to delete banker: %w", err)
	}
	return nil
}

func (b *BankerRepo) GetBanker(r *http.Request) ([]*models.Banker, error) {
	vars := mux.Vars(r)
	userId, ok := vars["user_id"]
	if !ok || userId == "" {
		return nil, errors.New("user_id cannot be empty")
	}

	query := database.NewQuery(b.db)
	bankers, err := query.GetBanker(userId)
	if err != nil {
		log.Println("Database error:", err)
		return nil, fmt.Errorf("failed to fetch bankers: %w", err)
	}

	return bankers, nil
}
