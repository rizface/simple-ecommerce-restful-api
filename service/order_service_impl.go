package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/repository"
	"sync"
	"time"
)

type orderService struct {
	db          *sql.DB
	validate    *validator.Validate
	productRepo repository.SellerProductRepository
	orderRepo   repository.OrderRepository
	invoiceRepo repository.InvoiceRepo
}

func NewOrderService(db *sql.DB, validate *validator.Validate, productRepo repository.SellerProductRepository, orderRepo repository.OrderRepository, invoiceRepo repository.InvoiceRepo) OrderService {
	return orderService{
		db:          db,
		validate:    validate,
		productRepo: productRepo,
		orderRepo:   orderRepo,
		invoiceRepo: invoiceRepo,
	}
}

func (o orderService) GetOrders() {
	panic("implement me")
}

func (o orderService) PostOrders(ctx context.Context, idCustomer int, request web.OrderRequest) string {
	err := o.validate.Struct(request)
	exception.PanicBadRequest(err)

	tx, err := o.db.Begin()
	exception.PanicIfInternalServerError(err)
	wg := sync.WaitGroup{}
	total := 0
	defer helper.CommitOrRollback(tx)
	invoice := time.Now().Format("2006-01-02") + helper.GenerateRandomString()

	wg.Add(len(request.Items))
	for i, v := range request.Items {
		product := o.productRepo.FindById(ctx, tx, v.IdProduct)
		exception.PanicNotFound(product.Id)
		order := web.Order{
			Invoice:    invoice,
			IdSeller:   product.IdSeller,
			IdCustomer: idCustomer,
			IdProduct:  product.Id,
			Jumlah:     v.Jumlah,
			Total:      v.Jumlah * product.HargaBarang,
			Alamat:     request.Alamat,
		}
		total += order.Total
		if i == len(request.Items) - 1 {
			o.invoiceRepo.PostInvoice(ctx, tx, invoice,total)
		}
		go o.orderRepo.PostOrders(ctx, tx, order, &wg)
	}
	wg.Wait()
	return "silahkan kirimkan bukti pembayaran masing - masing orderan secara terpisah, orderan akan dihapus secara otomatis jika tidak mengirimkan bukti pembayaran dalam 48 jam"
}
