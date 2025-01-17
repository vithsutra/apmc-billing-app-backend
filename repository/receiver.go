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

type ReceiverRepo struct {
	db *sql.DB
}

func NewReceiverRepo(db *sql.DB) *ReceiverRepo {
	return &ReceiverRepo{
		db: db,
	}
}

func (repo *ReceiverRepo) CreateReceiver(r *http.Request) error {
	var receiver models.Receiver

	if err := json.NewDecoder(r.Body).Decode(&receiver); err != nil {
		return err
	}

	if err := validator.Validate(receiver); err != nil {
		return err
	}

	receiver.ReceiverId = uuid.NewString()

	query := database.NewQuery(repo.db)

	if err := query.CreateReceiver(&receiver); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo *ReceiverRepo) DeleteReceiver(r *http.Request) error {
	vars := mux.Vars(r)

	receiverId := vars["receiver_id"]

	if receiverId == "" {
		return errors.New("Receiver id cannot empty")
	}

	query := database.NewQuery(repo.db)

	if err := query.DeleteReceiver(receiverId); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo *ReceiverRepo) GetReceivers(r *http.Request) ([]*models.Receiver, error) {
	vars := mux.Vars(r)

	userId := vars["user_id"]

	if userId == "" {
		return nil, errors.New("user id cannot be empty")
	}

	query := database.NewQuery(repo.db)

	receivers, err := query.GetReceivers(userId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return receivers, nil
}
