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
		SELECT user_id , user_password 
		FROM users
		WHERE user_email=$1
	`, email).Scan(&userId, &dbpassword)
	if err != nil {
		return "", err
	}
	err = utils.VerifyPassword(dbpassword, password)
	if err != nil {
		return "", err
	}
	ts, err := utils.GenerateJWTToken(userId)
	if err != nil {
		return "", err
	}
	return ts, nil
}

func (q *Query) Register(user models.User) (string, error) {
	user.UserId = uuid.NewString()
	pass, err := utils.EncryptPassword(user.UserPassword)
	if err != nil {
		return "", err
	}
	tk, err := utils.GenerateJWTToken(user.UserId)
	if err != nil {
		return "", err
	}
	_, err = q.db.Exec(`
		INSERT INTO users(
			user_id,
			user_name,
			user_email,
			user_password,
			user_address,
			user_phone,
			user_gstin,
			user_pan
		) VALUES(
		 	$1 , $2 , $3 , $4 , $5 , $6 , $7 , $8
		)
	`,
		user.UserId, user.UserName, user.UserEmail, pass,
		user.UserAddress, user.UserPhone, user.UserGSTIN, user.UserPAN,
	)
	if err != nil {
		return "", err
	}
	return tk, nil
}

func (q *Query) DeleteUser(userId string) error {
	query := `DELETE FROM users WHERE user_id=$1`
	_, err := q.db.Exec(query, userId)

	if err != nil {
		return err
	}

	return nil
}
