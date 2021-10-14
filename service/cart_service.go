package service

import (
	"context"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type CartService interface {
	PostItem(ctx context.Context, request web.CartRequest) string
	GetItems(ctx context.Context, idCustomer int) []domain.CartProduct
	UpdateItem(ctx context.Context, request web.CartRequest, idCart int) string
	DeleteItem(ctx context.Context, idCart int) string
}
