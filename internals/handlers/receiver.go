package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vsynclabs/billsoft/internals/models"
)

type ReceiverHandler struct {
	receiverRepo models.ReceiverInterface
}

func NewReceiverHandler(receiverRepo models.ReceiverInterface) *ReceiverHandler {
	return &ReceiverHandler{
		receiverRepo: receiverRepo,
	}
}

func (handler *ReceiverHandler) CreateReceiverHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := handler.receiverRepo.CreateReceiver(r); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "error occurred"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "created receiver successfully"})
}

func (handler *ReceiverHandler) DeleteReceiverHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := handler.receiverRepo.DeleteReceiver(r); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "error occurred"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "receiver deleted successfully"})
}

func (handler *ReceiverHandler) GetReceiversHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	receivers, err := handler.receiverRepo.GetReceivers(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "error occurrred"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*models.Receiver{"receiver_details": receivers})
}
