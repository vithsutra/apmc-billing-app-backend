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
	db               *sql.DB
	emailServiceRepo models.EmailServiceInterface
	otpCache         map[string]otpData
}

type otpData struct {
	OTP       string
	ExpiresAt time.Time
}

func NewUserRepo(db *sql.DB, emailServiceRepo models.EmailServiceInterface) *UserRepo {
	return &UserRepo{
		db:               db,
		emailServiceRepo: emailServiceRepo,
		otpCache:         make(map[string]otpData),
	}
}

func (ur *UserRepo) Login(r *http.Request) (string, error) {
	var userLoginRequest models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&userLoginRequest); err != nil {
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
		return "", err
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
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

func (ur *UserRepo) UserForgotPassword(r *http.Request) (int32, error) {
	var forgotPasswordRequest models.ForgotPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&forgotPasswordRequest); err != nil {
		return 400, errors.New("invalid JSON request body")
	}

	validate := validator.New()
	if err := validate.Struct(forgotPasswordRequest); err != nil {
		return 400, errors.New("invalid request format")
	}

	query := database.NewQuery(ur.db)
	emailExists, err := query.CheckUserEmailsExists(forgotPasswordRequest.UserEmail)
	if err != nil {
		log.Println(err)
		return 500, errors.New("internal server error")

	}
	if !emailExists {
		return 400, errors.New("email does not exist")
	}

	otp, err := utils.GenerateOTP()
	log.Println(otp)
	if err != nil {
		return 500, errors.New("failed to generate OTP")
	}

	expireTime := time.Now().Add(5 * time.Minute)

	if err := query.StoreUserOtp(forgotPasswordRequest.UserEmail, otp, expireTime); err != nil {
		return 500, errors.New("failed to store OTP")
	}

	emailBody := &models.UserOtpEmailFormat{
		To:        forgotPasswordRequest.UserEmail,
		Subject:   "Verification Code to Reset Password",
		EmailType: "otp",
		Data: map[string]string{
			"otp":         otp,
			"expire_time": "5",
		},
	}

	jsonBytes, err := json.Marshal(emailBody)
	if err != nil {
		return 500, errors.New("error generating email body")
	}

	if err := ur.emailServiceRepo.SendEmail(jsonBytes); err != nil {
		return 500, errors.New("failed to send email")
	}

	go func() {
		time.Sleep(5 * time.Minute)
		_ = query.DeleteUserOtp((forgotPasswordRequest.UserEmail), otp)
	}()

	return 200, nil
}

func (ur *UserRepo) ValidateOTP(r *http.Request) (string, error) {
	var otpReq models.OTPValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&otpReq); err != nil {
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
	if err := query.UpdateUserPassword(emailData.OTP, resetReq.NewPassword); err != nil {
		return err
	}

	delete(ur.otpCache, resetReq.TokenId)
	return nil
}
