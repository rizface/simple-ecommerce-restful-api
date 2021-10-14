package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/repository"
)

type cartServiceImpl struct {
	db          *sql.DB
	validate    *validator.Validate
	cartRepo    repository.CartRepository
	productRepo repository.CustomerProduct
}

func NewCartService(db *sql.DB, validate *validator.Validate, cartRepo repository.CartRepository, productRepo repository.CustomerProduct) CartService {
	return cartServiceImpl{
		db:          db,
		validate:    validate,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (c cartServiceImpl) PostItem(ctx context.Context, request web.CartRequest) string {
	err := c.validate.Struct(request)
	exception.PanicBadRequest(err)

	tx,err := c.db.Begin()
	exception.PanicIfInternalServerError(err)

	defer helper.CommitOrRollback(tx)
	product := c.productRepo.GetDetail(ctx,tx,request.IdProduct)
	exception.PanicNotFound(product.Id)

	if product.StokBarang < request.Jumlah {
		exception.PanicBadRequest(errors.New("jumlah yang kamu masukan tidak valid"))
	}

	var success bool
	cartExist := c.cartRepo.GetItemsByIdCustomerandProduct(ctx,tx,request.IdCustomer,request.IdProduct)
	if cartExist.IdCart > 0 {
		return  c.UpdateItem(ctx,request,cartExist.IdCart)
	} else {
		success = c.cartRepo.PostItem(ctx,tx,request)
		if success {
			return "success"
		}
		return "failed"
	}

}

func (c cartServiceImpl) GetItems(ctx context.Context, idCustomer int)[]domain.CartProduct {
	tx,err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	items := c.cartRepo.GetItems(ctx,tx,idCustomer)
	return items
}

func (c cartServiceImpl) UpdateItem(ctx context.Context,request web.CartRequest, idCart int) string {
	tx,err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	exist := c.cartRepo.GetCart(ctx,tx,idCart)
	exception.PanicNotFound(exist.IdCart)
	existProduct := c.productRepo.GetDetail(ctx,tx,exist.IdProduct)
	exception.PanicNotFound(existProduct.Id)
	if request.Jumlah > existProduct.StokBarang {
		exception.PanicBadRequest(errors.New("jumlah yang kamu masukan tidak valid"))
	} else if request.Jumlah < 1 {
		return c.DeleteItem(ctx,exist.IdCart)
	}
	success := c.cartRepo.UpdateItem(ctx,tx,request,exist.IdCart)
	if success {
		return "success"
	}
	return "failed"
}

func (c cartServiceImpl) DeleteItem(ctx context.Context, idCart int) string {
	tx,err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)

	exist := c.cartRepo.GetCart(ctx,tx,idCart)
	exception.PanicNotFound(exist.IdCart)
	success :=  c.cartRepo.DeleteCart(ctx,tx,exist.IdCart)
	if success {
		return "success"
	}
	return "failed"
}

