package database

import (
	"github.com/vsynclabs/billsoft/internals/models"
)

func (q *Query) CreateInvoice(invoice *models.Invoice) error {
	query := `INSERT INTO invoice (
					invoice_id,
					name,
					payment_status,
					user_id,
					billed_id,
					shipped_id,
					invoice_date,
					supply_date
				) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := q.db.Exec(
		query,
		invoice.InvoiceId,
		invoice.Name,
		invoice.PaymentStatus,
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

func (q *Query) GetInvoices(userId string) ([]*models.InvoiceResponse, error) {
	query := `SELECT invoice_id,name,payment_status FROM invoice WHERE user_id = $1`

	rows, err := q.db.Query(query, userId)

	if err != nil {
		return nil, err
	}

	var invoices []*models.InvoiceResponse

	for rows.Next() {
		var invoice models.InvoiceResponse

		err := rows.Scan(&invoice.InvoiceId, &invoice.Name, &invoice.PaymentStatus)

		if err != nil {
			return nil, err
		}

		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}
