package database

import "github.com/vsynclabs/billsoft/internals/models"

func (q *Query) CreateProduct(product *models.Product) error {
	query := `INSERT INTO product(
		product_id,
		product_name,
		product_hsn,
		product_quantity,
		product_unit,
		product_rate,
		product_total,
		invoice_id
	)VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := q.db.Exec(query,
		product.ProductId,
		product.ProductName,
		product.ProductHsn,
		product.ProductQty,
		product.ProductUnit,
		product.ProductRate,
		product.Total,
		product.InvoiceId,
	)
	return err
}

func (q *Query) DeleteProduct(productId string) error {
	query := `DELETE FROM product WHERE product_id =$1`
	_, err := q.db.Exec(query, productId)
	return err
}

func (q *Query) GetProduct(invoiceId string) ([]*models.Product, error) {
	query := `SELECT 
				product_id,
				product_name,
				product_hsn,
				product_quantity,
				product_unit,
				product_rate,
				product_total
			FROM product WHERE invoice_id=$1`

	rows, err := q.db.Query(query, invoiceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		var product models.Product

		if err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.ProductHsn,
			&product.ProductQty,
			&product.ProductUnit,
			&product.ProductRate,
			&product.Total,
		); err != nil {
			return nil, err
		}
		products = append(products, &product)

	}

	return products, nil
}
