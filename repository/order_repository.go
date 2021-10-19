package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/model/web"
	"sync"
)

type OrderRepository interface {
	GetOrders()
	PostOrders(ctx context.Context, tx *sql.Tx, request web.Order, wg *sync.WaitGroup)
}
