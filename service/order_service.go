package service

import (
	"context"
	"simple-ecommerce-rest-api/model/web"
)

type OrderService interface {
	GetOrders()
	PostOrders(ctx context.Context, idCustomer int, request web.OrderRequest) string
}
