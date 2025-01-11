package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userId string) (string , error) {
	var secretKey []byte = []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenString , err := claims.SignedString(secretKey)
	if err != nil {
		return "" , err
	}
	return tokenString , nil
}