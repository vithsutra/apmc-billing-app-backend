package database

import (
	"github.com/google/uuid"
	"github.com/vsynclabs/billsoft/internals/models"
	"github.com/vsynclabs/billsoft/pkg/utils"
)

func (q *Query) Login(email, password string) (string, error) {
	var dbpassword string
	var userId string

	err := q.db.QueryRow(`
		SELECT user_id, user_password
		FROM users
		WHERE user_email = $1
	`, email).Scan(&userId, &dbpassword)
	if err != nil {
		return "", err
	}

	err = utils.VerifyPassword(dbpassword, password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWTToken(userId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (q *Query) Register(user models.User) (string, error) {
	user.UserId = uuid.NewString()

	hashedPassword, err := utils.EncryptPassword(user.UserPassword)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWTToken(user.UserId)
	if err != nil {
		return "", err
	}

	_, err = q.db.Exec(`
		INSERT INTO users (
			user_id,
			user_name,
			user_email,
			user_password,
			user_phone
		) VALUES ($1, $2, $3, $4, $5)
	`, user.UserId, user.UserName, user.UserEmail, hashedPassword, user.UserPhone)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (q *Query) DeleteUser(userId string) error {
	_, err := q.db.Exec(`DELETE FROM users WHERE user_id = $1`, userId)
	return err
}

func (q *Query) UpdateUserPassword(email string, newPassword string) error {
	hashedPassword, err := utils.EncryptPassword(newPassword)
	if err != nil {
		return err
	}

	_, err = q.db.Exec(`
		UPDATE users
		SET user_password = $1
		WHERE user_email = $2
	`, hashedPassword, email)

	return err
}
