package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type CustomerRepository interface {
	RegisterCustomer(ctx context.Context, tx *sql.Tx, request web.RequestCustomer) int64
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) domain.Customers
	UpdateConfirmed(ctx context.Context, tx *sql.Tx, idCustomer int) bool
}
