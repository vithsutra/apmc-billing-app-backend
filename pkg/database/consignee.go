package database

import "github.com/vsynclabs/billsoft/internals/models"

func (q *Query) CreateConsignee(consignee *models.Consignee) error {
	query := `INSERT INTO shipped(
		shipped_id,
		shipped_name,
		shipped_address,
		shipped_gstin,
		shipped_mobile,
		shipped_state,
		shipped_state_code,
		user_id
	
	)VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := q.db.Exec(query,
		consignee.ConsigneeId,
		consignee.ConsigneeName,
		consignee.ConsigneeAddress,
		consignee.ConsigneeGstin,
		consignee.ConsigneePhoneNumber,
		consignee.ConsigneeState,
		consignee.ConsigneeStateCode,
		consignee.UserId,
	)
	return err

}

func (q *Query) DeleteConsignee(consigneeId string) error {
	query := `DELETE FROM shipped WHERE shipped_id=$1`
	_, err := q.db.Exec(query, consigneeId)
	return err
}

func (q *Query) GetConsignee(userId string) ([]*models.Consignee, error) {
	query := `SELECT 
		shipped_id,
		shipped_name,
		shipped_address,
		shipped_gstin,
		shipped_mobile,
		shipped_state,
		shipped_state_code
	FROM shipped WHERE user_id = $1`

	rows, err := q.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consignees []*models.Consignee
	for rows.Next() {
		var consignee models.Consignee
		if err := rows.Scan(
			&consignee.ConsigneeId,
			&consignee.ConsigneeName,
			&consignee.ConsigneeAddress,
			&consignee.ConsigneeGstin,
			&consignee.ConsigneePhoneNumber,
			&consignee.ConsigneeState,
			&consignee.ConsigneeStateCode,
		); err != nil {
			return nil, err
		}
		consignees = append(consignees, &consignee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return consignees, nil
}
