package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/model/web"
	"sync"
)

type orderRepository struct {}

func NewOrderRepository() OrderRepository {
	return orderRepository{}
}

func (o orderRepository) GetOrders() {
	panic("implement me")
}

func (o orderRepository) PostOrders(ctx context.Context, tx *sql.Tx, request web.Order, wg *sync.WaitGroup) {
	defer wg.Done()
	sql := "INSERT INTO orders(invoice,id_seller,id_customer,id_product,jumlah,total,alamat) VALUES(?,?,?,?,?,?,?)"
	result,err := tx.ExecContext(ctx,sql,request.Invoice,request.IdSeller,request.IdCustomer,request.IdProduct,request.Jumlah,request.Total,request.Alamat)

	exception.PanicIfInternalServerError(err)
	affected,_ := result.RowsAffected()
	if affected == 0 {
		exception.PanicIfInternalServerError(err)
	}
}
