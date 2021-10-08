package service

import (
	"context"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type SellerService interface {
	Register(ctx context.Context, request web.RequestSellerRegister) domain.Seller
}

