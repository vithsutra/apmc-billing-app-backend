package database

import (
	"github.com/vsynclabs/billsoft/internals/models"
)

func (q *Query) CreateBanker(banker *models.Banker) error {
	query := `INSERT INTO banker (
				bank_id,
				bank_name,
				bank_branch,
				bank_account_number,
				bank_ifsc_code,
				bank_holder_name,
				user_id
				) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := q.db.Exec(
		query,
		banker.BankId,
		banker.BankName,
		banker.BankBranch,
		banker.BankAccountNumber,
		banker.BankIfscCode,
		banker.BankHolderName,
		banker.UserId,
	)
	return err
}

func (q *Query) DeleteBanker(BankId string) error {
	query := `DELETE FROM banker WHERE bank_id=$1`
	_, err := q.db.Exec(query, BankId)
	return err
}

func (q *Query) GetBanker(userId string) ([]*models.Banker, error) {
	query := `SELECT bank_id, bank_name, bank_branch, bank_account_number, bank_ifsc_code, bank_holder_name, user_id FROM banker WHERE user_id=$1`
	rows, err := q.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bankers []*models.Banker
	for rows.Next() {
		var banker models.Banker
		err := rows.Scan(
			&banker.BankId,
			&banker.BankName,
			&banker.BankBranch,
			&banker.BankAccountNumber,
			&banker.BankIfscCode,
			&banker.BankHolderName,
			&banker.UserId,
		)
		if err != nil {
			return nil, err
		}
		bankers = append(bankers, &banker)
	}

	return bankers, nil
}
