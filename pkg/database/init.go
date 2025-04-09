package database

import (
	"log"
)

func (q *Query) InitilizeDatabase() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users(
			user_id VARCHAR(100) PRIMARY KEY,
			user_name VARCHAR(100) NOT NULL,
			user_phone VARCHAR(100) NOT NULL,
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

		`CREATE TABLE IF NOT EXISTS biller (
			biller_id VARCHAR(100) PRIMARY KEY,
			biller_name VARCHAR(100) NOT NULL,
			biller_address VARCHAR(500) NOT NULL,
			biller_mobile VARCHAR(100) NOT NULL,
			biller_gstin VARCHAR(100) NOT NULL,
			biller_pan VARCHAR(100) NOT NULL,
			biller_mail VARCHAR(100) NOT NULL,
			biller_companylogo VARCHAR(100) NOT NULL DEFAULT 'PENDING',
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
			invoice_name VARCHAR(100) NOT NULL,
			invoice_payment_status BOOLEAN NOT NULL,
			invoice_reverse_charge VARCHAR(100)NOT NULL,
			invoice_number SERIAL UNIQUE,
			invoice_date VARCHAR(100) NOT NULL,
			invoice_state VARCHAR(100) NOT NULL,
			invoice_state_code VARCHAR(100) NOT NULL,
			invoice_challan_number VARCHAR(100) NOT NULL,
			invoice_vehicle_number VARCHAR(100) NOT NULL,
			invoice_date_of_supply VARCHAR(100),
			invoice_place_of_supply VARCHAR(100),
			invoice_gst VARCHAR(100) NOT NULL,
			user_id VARCHAR(100) NOT NULL,
			billed_id VARCHAR(100) NOT NULL,
			shipped_id VARCHAR(100) NOT NULL,
			biller_id VARCHAR(100) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (billed_id) REFERENCES billed(billed_id) ON DELETE CASCADE,
			FOREIGN KEY (shipped_id) REFERENCES shipped(shipped_id) ON DELETE CASCADE,
			FOREIGN KEY (biller_id) REFERENCES biller(biller_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS product(
			product_id VARCHAR(100) PRIMARY KEY,
			product_name VARCHAR(100) NOT NULL,
			product_hsn VARCHAR(100) NOT NULL,
			product_quantity VARCHAR(100) NOT NULL,
			product_unit VARCHAR(100) NOT NULL,
			product_rate VARCHAR(100) NOT NULL,
			product_total VARCHAR(100) NOT NULL,
			invoice_id VARCHAR(100) NOT NULL,
			FOREIGN KEY (invoice_id) REFERENCES invoice(invoice_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS banker(
			bank_id VARCHAR(100) PRIMARY KEY,
			bank_name VARCHAR(100) NOT NULL,
			bank_branch VARCHAR(100) NOT NULL,
			bank_account_number VARCHAR(100) NOT NULL,
			bank_ifsc_code VARCHAR(100) NOT NULL,
			bank_holder_name VARCHAR(100) NOT NULL,
			user_id VARCHAR(100) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS user_otps (
			email VARCHAR(255) NOT NULL,
			otp VARCHAR(255) NOT NULL,
			expire_time TIMESTAMPTZ NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			FOREIGN KEY (email) REFERENCES users(user_email) ON DELETE CASCADE

		)
			`,
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
			log.Printf("Failed Query: %s\nError: %v", query, err)
			return err
		}
	}
	log.Println("Database initilized successfully")
	return nil
}
