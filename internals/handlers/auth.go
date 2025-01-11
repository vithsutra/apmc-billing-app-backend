package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vsynclabs/billsoft/internals/models"
)

type UserHandler struct{
	UserRepo models.UserInterface
}

func NewUserHandler(UserRepo models.UserInterface) *UserHandler {
	return &UserHandler{
		UserRepo,
	}
}

func(uh *UserHandler) LoginHandler(w http.ResponseWriter , r *http.Request) {
	res , err := uh.UserRepo.Login(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message":err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token":res})
}

func(uh *UserHandler) RegisterHandler(w http.ResponseWriter , r *http.Request) {
	res , err := uh.UserRepo.Register(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message":err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token":res})
}