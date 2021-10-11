package service

import (
	"context"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type SellerProductService interface {
	GetProducts(ctx context.Context, idSeller int) []domain.Products
	GetDetailProduct(ctx context.Context, idProduct int) domain.Products
	PostProduct(ctx context.Context, idSeller int, request web.ProductRequest) domain.Products
	DeleteProduct(ctx context.Context, idProduct int) bool
	UpdateProduct(ctx context.Context, idProduct int, idSeller int, request web.ProductRequest) string
}