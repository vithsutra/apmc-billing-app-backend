package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/database"
	"github.com/vsynclabs/billsoft/pkg/utils"
)

type UserRepo struct {
	db       *sql.DB
	otpCache map[string]otpData
}

type otpData struct {
	OTP       string
	ExpiresAt time.Time
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db:       db,
		otpCache: make(map[string]otpData),
	}
}

func (ur *UserRepo) Login(r *http.Request) (string, error) {
	var userLoginRequest models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&userLoginRequest); err != nil {
		log.Println(err)
		return "", err
	}

	validate := validator.New()
	if err := validate.Struct(userLoginRequest); err != nil {
		return "", err
	}

	query := database.NewQuery(ur.db)
	return query.Login(userLoginRequest.UserEmail, userLoginRequest.UserPassword)
}

func (ur *UserRepo) Register(r *http.Request) (string, error) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		return "", err
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Println(err)
		return "", err
	}
	query := database.NewQuery(ur.db)
	return query.Register(user)
}

func (ur *UserRepo) DeleteUser(r *http.Request) error {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	if userId == "" {
		return errors.New("user id is required")
	}
	query := database.NewQuery(ur.db)
	return query.DeleteUser(userId)
}

func (ur *UserRepo) ForgotPassword(r *http.Request) error {
	var forgotReq models.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&forgotReq); err != nil {
		log.Println(err)
		return err
	}

	validate := validator.New()
	if err := validate.Struct(forgotReq); err != nil {
		return err
	}

	otp := utils.GenerateOTP(6)
	ur.otpCache[forgotReq.UserEmail] = otpData{
		OTP:       otp,
		ExpiresAt: time.Now().Add(1 * time.Minute),
	}

	return utils.SendResetTokenMail(forgotReq.UserEmail, otp)
}

func (ur *UserRepo) ValidateOTP(r *http.Request) (string, error) {
	var otpReq models.OTPValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&otpReq); err != nil {
		log.Println(err)
		return "", err
	}

	validate := validator.New()
	if err := validate.Struct(otpReq); err != nil {
		return "", err
	}

	storedOTP, ok := ur.otpCache[otpReq.UserEmail]
	if !ok || time.Now().After(storedOTP.ExpiresAt) {
		return "", fmt.Errorf("OTP expired or not found")
	}
	if otpReq.OTP != storedOTP.OTP {
		return "", fmt.Errorf("Invalid OTP")
	}

	tokenId := uuid.NewString()
	delete(ur.otpCache, otpReq.UserEmail)

	ur.otpCache[tokenId] = otpData{
		OTP: otpReq.UserEmail,
	}

	return tokenId, nil
}

func (ur *UserRepo) ResetPassword(r *http.Request) error {
	var resetReq models.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&resetReq); err != nil {
		log.Println(err)
		return err
	}

	validate := validator.New()
	if err := validate.Struct(resetReq); err != nil {
		return err
	}

	emailData, ok := ur.otpCache[resetReq.TokenId]
	if !ok {
		return fmt.Errorf("invalid or expired token")
	}

	query := database.NewQuery(ur.db)
	err := query.UpdateUserPassword(emailData.OTP, resetReq.NewPassword)
	if err != nil {
		return err
	}

	delete(ur.otpCache, resetReq.TokenId)
	return nil
}
