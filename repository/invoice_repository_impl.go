package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/app/exception"
)

type invoiceRepo struct{}

func NewInvoiceRepo() InvoiceRepo {
	return invoiceRepo{}
}

func (i invoiceRepo) PostInvoice(ctx context.Context, tx *sql.Tx, invoice string,total int) {
	sql := "INSERT INTO invoice(invoice,total) VALUES(?,?)"
	_, err := tx.ExecContext(ctx,sql,invoice,total)
	exception.PanicIfInternalServerError(err)
}
