package models

import "net/http"

type User struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserAddress  string `json:"user_address"`
	UserPhone    string `json:"user_phone"`
	UserGSTIN    string `json:"user_gstin"`
	UserPAN      string `json:"user_pan"`
}

type UserInterface interface {
	Login(*http.Request) (string, error)
	Register(*http.Request) (string, error)
}
