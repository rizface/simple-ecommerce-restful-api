package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type CartRepository interface{
	UpdateItem(ctx context.Context, tx *sql.Tx, request web.CartRequest, idCart int) bool
	PostItem(ctx context.Context,tx *sql.Tx,request web.CartRequest) bool
	GetItems(ctx context.Context, tx *sql.Tx, idCustomer int) []domain.CartProduct
	GetCart(ctx context.Context, tx *sql.Tx, idCart int) domain.CartProduct
	GetItemsByIdCustomerandProduct(ctx context.Context, tx *sql.Tx, idCustomer int, idProduct int) domain.CartProduct
	DeleteCart(ctx context.Context, tx *sql.Tx, idCart int) bool

}
