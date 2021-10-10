package service

import (
	"context"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type SellerProductService interface {
	GetProducts(ctx context.Context, idSeller int) []domain.Products
	PostProduct(ctx context.Context, idSeller int, request web.NewProduct) domain.Products
}