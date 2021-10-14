package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type CustomerRepository interface {
	RegisterCustomer(ctx context.Context, tx *sql.Tx, request web.RequestCustomer) bool
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) domain.Customers
}
