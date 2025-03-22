package database

import (
	"github.com/vsynclabs/billsoft/internals/models"
)

func (q *Query) CreateBiller(biller *models.Biller) error {
	if biller.BillerCompanyLogo == "" {
		biller.BillerCompanyLogo = "PENDING"
	}

	query := `INSERT INTO biller(
        biller_id,
        biller_name,
        biller_address,
        biller_mobile,
        biller_gstin,
        biller_pan,
        biller_companylogo,
        biller_mail,
        user_id
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := q.db.Exec(query,
		biller.BillerId,
		biller.BillerName,
		biller.BillerAddress,
		biller.BillerMobile,
		biller.BillerGstin,
		biller.BillerPan,
		biller.BillerCompanyLogo,
		biller.BillerMail,
		biller.UserId,
	)
	return err
}

func (q *Query) DeleteBiller(billerId string) error {
	query := `DELETE FROM biller WHERE biller_id = $1`
	_, err := q.db.Exec(query, billerId)
	return err
}

func (q *Query) GetBiller(userId string) ([]*models.Biller, error) {
	query := `SELECT biller_id, biller_name, biller_address, biller_mobile, biller_gstin, biller_pan, biller_companylogo, biller_mail, user_id 
	          FROM biller WHERE user_id = $1`
	rows, err := q.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var billers []*models.Biller
	for rows.Next() {
		var biller models.Biller
		err := rows.Scan(
			&biller.BillerId,
			&biller.BillerName,
			&biller.BillerAddress,
			&biller.BillerMobile,
			&biller.BillerGstin,
			&biller.BillerPan,
			&biller.BillerCompanyLogo,
			&biller.BillerMail,
			&biller.UserId,
		)
		if err != nil {
			return nil, err
		}
		billers = append(billers, &biller)
	}
	return billers, nil
}
