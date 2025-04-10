package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vsynclabs/billsoft/internals/models"
)

type UserHandler struct {
	UserRepo models.UserDatabaseInterface
}

func NewUserHandler(UserRepo models.UserDatabaseInterface) *UserHandler {
	return &UserHandler{
		UserRepo,
	}
}

func (uh *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token, err := uh.UserRepo.Login(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "success",
		Message: "login successful",
		Data:    map[string]string{"token": token},
	})
}

func (uh *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token, err := uh.UserRepo.Register(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "success",
		Message: "registration successful",
		Data:    map[string]string{"token": token},
	})
}

func (uh *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := uh.UserRepo.DeleteUser(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "success",
		Message: "user deleted successfully",
	})
}

func (uh *UserHandler) ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusCode, err := uh.UserRepo.UserForgotPassword(r)
	if err != nil {
		w.WriteHeader(int(statusCode))
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}
	w.WriteHeader(int(statusCode))
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "success",
		Message: "OTP sent successfully",
	})
}

func (uh *UserHandler) ValidateOTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenId, err := uh.UserRepo.ValidateOTP(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "success",
		Message: "OTP validated successfully",
		Data:    models.TokenResponse{TokenId: tokenId},
	})
}

func (uh *UserHandler) ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := uh.UserRepo.ResetPassword(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "success",
		Message: "password reset successful",
	})
}
