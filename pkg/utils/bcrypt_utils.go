package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(originalPassword, userPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(userPassword)); err != nil {
		return err
	}
	return nil
}
