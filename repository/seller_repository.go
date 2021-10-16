package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type SellerRepository interface {
	Register(ctx context.Context, tx *sql.Tx, request web.RequestSeller) int64
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) domain.Seller
	FindByName(ctx context.Context, tx *sql.Tx, name string) domain.Seller
	Confirm(ctx context.Context,tx *sql.Tx, idSeller int) bool
}
