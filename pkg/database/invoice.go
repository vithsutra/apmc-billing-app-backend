package database

import "github.com/vsynclabs/billsoft/internals/models"

func (q *Query) CreateInvoice(invoice *models.Invoice) error {
	query := `INSERT INTO invoice (
					invoice_id,
					user_id,
					billed_id,
					shipped_id,
					invoice_date,
					supply_date,
				) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := q.db.Exec(
		query,
		invoice.InvoiceId,
		invoice.UserId,
		invoice.ReceiverId,
		invoice.ConsigneeId,
		invoice.InvoiceDate,
		invoice.SupplyDate,
	)

	return err
}

func (q *Query) DeleteInvoice(invoiceId string) error {
	query := `DELETE FROM invoice WHERE invoice_id=$1`

	_, err := q.db.Exec(query, invoiceId)

	return err
}
