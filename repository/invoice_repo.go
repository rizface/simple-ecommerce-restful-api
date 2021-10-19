package repository

import (
	"context"
	"database/sql"
)

type InvoiceRepo interface {
	PostInvoice(ctx context.Context, tx *sql.Tx, invoice string, total int)
}
