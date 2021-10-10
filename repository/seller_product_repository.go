package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
	"sync"
)

type SellerProductRepository interface {
	GetProducts(ctx context.Context, tx *sql.Tx, idSeller int) []domain.Products
	PostProduct(ctx context.Context, tx *sql.Tx, idSeller int, request web.NewProduct) domain.Products
	PostProductImages(ctx context.Context, tx *sql.Tx, idProduct int, link string, wg *sync.WaitGroup)
}
