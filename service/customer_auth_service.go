package service

import (
	"context"
	"simple-ecommerce-rest-api/model/web"
)

type CustomerAuthService interface {
	RegisterCustomer(ctx context.Context, request web.RequestCustomer) bool
	LoginCustomer(ctx context.Context, request web.RequestCustomer) string
}
