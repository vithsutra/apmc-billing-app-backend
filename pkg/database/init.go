package database

import (
	"log"
)

func (q *Query) InitilizeDatabase() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users(
			user_id VARCHAR(100) PRIMARY KEY,
			user_name VARCHAR(100) NOT NULL,
			user_address VARCHAR(500) NOT NULL,
			user_phone VARCHAR(100) NOT NULL,
			user_gstin VARCHAR(100) NOT NULL,
			user_pan VARCHAR(100) NOT NULL,
			user_email VARCHAR(200) NOT NULL UNIQUE,
			user_password VARCHAR(100) NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS billed(
			billed_id VARCHAR(100) PRIMARY KEY,
			billed_name VARCHAR(100) NOT NULL,
			billed_address VARCHAR(500) NOT NULL,
			billed_gstin VARCHAR(100) NOT NULL,
			billed_state VARCHAR(100) NOT NULL,
			billed_state_code VARCHAR(100) NOT NULL,
			user_id VARCHAR(100) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS shipped(
			shipped_id VARCHAR(100) PRIMARY KEY,
			shipped_name VARCHAR(100) NOT NULL,
			shipped_address VARCHAR(500) NOT NULL,
			shipped_gstin VARCHAR(100) NOT NULL,
			shipped_mobile VARCHAR(100) NOT NULL,
			shipped_state VARCHAR(100) NOT NULL,
			shipped_state_code VARCHAR(100) NOT NULL,
			user_id VARCHAR(100) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS invoice(
			invoice_id VARCHAR(100) PRIMARY KEY,
			user_id VARCHAR(100) NOT NULL,
			billed_id VARCHAR(100) NOT NULL,
			shipped_id VARCHAR(100) NOT NULL,
			invoice_date VARCHAR(100),
			supply_date VARCHAR(100),
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (billed_id) REFERENCES billed(billed_id) ON DELETE CASCADE,
			FOREIGN KEY (shipped_id) REFERENCES shipped(shipped_id) ON DELETE CASCADE
		)`,
	}
	tx, err := q.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	for _, query := range queries {
		_, err = tx.Exec(query)
		if err != nil {
			return err
		}
	}
	log.Println("Database initilized successfully")
	return nil
}
