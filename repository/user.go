package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
	validator "gopkg.in/validator.v2"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db,
	}
}

func (ur *UserRepo) Login(r *http.Request) (string, error) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return "", err
	}
	if err := validator.Validate(user); err != nil {
		return "", err
	}
	query := database.NewQuery(ur.db)
	tk, err := query.Login(user.UserEmail, user.UserPassword)
	if err != nil {
		return "", err
	}
	return tk, nil
}

func (ur *UserRepo) Register(r *http.Request) (string, error) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return "", err
	}
	if err := validator.Validate(user); err != nil {
		return "", err
	}
	query := database.NewQuery(ur.db)
	tk, err := query.Register(user)
	if err != nil {
		return "", err
	}
	return tk, nil
}

func (ur *UserRepo) DeleteUser(r *http.Request) error {
	vars := mux.Vars(r)

	userId := vars["user_id"]

	if userId == "" {
		return errors.New("user id is required")
	}

	query := database.NewQuery(ur.db)

	if err := query.DeleteUser(userId); err != nil {
		return err
	}

	return nil
}
