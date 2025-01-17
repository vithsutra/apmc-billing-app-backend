package database

import "github.com/vsynclabs/billsoft/internals/models"

func (q *Query) CreateReceiver(receiver *models.Receiver) error {
	query := `INSERT INTO billed (
				billed_id,
				billed_name,
				billed_address,
				billed_gstin,
				billed_state,
				billed_state_code,
				user_id
			) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	_, err := q.db.Exec(query,
		receiver.ReceiverId,
		receiver.ReceiverName,
		receiver.ReceiverAddress,
		receiver.ReceiverGstin,
		receiver.ReceiverState,
		receiver.ReceiverStateCode,
		receiver.UserId,
	)

	return err
}

func (q *Query) DeleteReceiver(receiverId string) error {
	query := `DELETE FROM billed WHERE billed_id=$1`
	_, err := q.db.Exec(query, receiverId)
	return err
}
