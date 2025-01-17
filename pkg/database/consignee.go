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
