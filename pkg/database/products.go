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
		invoice_id
	)VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := q.db.Exec(query,
		product.ProductId,
		product.ProductName,
		product.ProductHsn,
		product.ProductQty,
		product.ProductUnit,
		product.ProductRate,
		product.InvoiceId,
	)
	return err
}

func (q *Query) DeleteProduct(productId string) error {
	query := `DELETE FROM product WHERE product_id =$1`
	_, err := q.db.Exec(query, productId)
	return err
}
