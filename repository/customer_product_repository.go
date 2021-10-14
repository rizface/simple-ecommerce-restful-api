package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/model/domain"
)

type CustomerProduct interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.Products
	GetDetail(ctx context.Context, tx *sql.Tx, idProduct int) domain.Products
}
