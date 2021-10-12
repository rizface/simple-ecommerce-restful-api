package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/model/domain"
)

type ProductImagesRepository interface {
	GetByProductId(ctx context.Context, tx *sql.Tx, idProduct int) []domain.ProductImages
}
