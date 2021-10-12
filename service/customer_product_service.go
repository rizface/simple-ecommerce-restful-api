package service

import (
	"context"
	"simple-ecommerce-rest-api/model/domain"
)

type CustomerProductService interface {
	Get(ctx context.Context) []domain.Products
	GetDetail(ctx context.Context, idProduct int) domain.Products
}