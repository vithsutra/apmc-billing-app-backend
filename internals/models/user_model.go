package models

import "net/http"

type UserLoginRequest struct {
	UserEmail    string `json:"user_email" validate:"required,email,max=50"`
	UserPassword string `json:"user_password" validate:"required,max=50"`
}
type User struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name" validate:"required,max=50"`
	UserEmail    string `json:"user_email" validate:"required,email,max=50"`
	UserPassword string `json:"user_password" validate:"required,max=50"`
	UserPhone    string `json:"user_phone" validate:"required,max=50"`
}

type ForgotPasswordRequest struct {
	UserEmail string `json:"user_email" validate:"required,email,max=50"`
}
type UserOtpEmailFormat struct {
	To        string            `json:"to"`
	Subject   string            `json:"subject"`
	EmailType string            `json:"email_type"`
	Data      map[string]string `json:"data"`
}

type OTPValidationRequest struct {
	UserEmail string `json:"user_email" validate:"required,email,max=50"`
	OTP       string `json:"otp" validate:"required,len=6,numeric"`
}

type TokenResponse struct {
	TokenId string `json:"token_id"`
}

type ResetPasswordRequest struct {
	TokenId         string `json:"token_id" validate:"required,uuid"`
	NewPassword     string `json:"new_password" validate:"required,max=50"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}

type UserDatabaseInterface interface {
	Login(r *http.Request) (string, error)
	Register(r *http.Request) (string, error)
	DeleteUser(r *http.Request) error
	ResetPassword(r *http.Request) error
	ValidateOTP(r *http.Request) (string, error)
	UserForgotPassword(r *http.Request) (int32, error)
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type EmailServiceInterface interface {
	SendEmail(data []byte) error
}
