package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/repository"
	"sync"
)

type sellerProductServiceImpl struct {
	db *sql.DB
	validate *validator.Validate
	sellerProduct repository.SellerProductRepository
}

func NewSellerProductServiceImpl(db *sql.DB, validate *validator.Validate, sellerProduct repository.SellerProductRepository) SellerProductService{
	return &sellerProductServiceImpl{
		db:            db,
		validate:      validate,
		sellerProduct: sellerProduct,
	}
}

func (s *sellerProductServiceImpl) GetProducts(ctx context.Context, idSeller int)[]domain.Products {
	tx,err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	products := s.sellerProduct.GetProducts(ctx,tx,idSeller)
	return products
}

func (s *sellerProductServiceImpl) PostProduct(ctx context.Context, idSeller int, request web.NewProduct) domain.Products {
	err := s.validate.Struct(request)
	exception.PanicBadRequest(err)

	images := helper.UploadProductImages(request.Gambar)
	request.Gambar = images

	tx,err := s.db.Begin()
	wg := sync.WaitGroup{}
	exception.PanicBadRequest(err)
	defer helper.CommitOrRollback(tx)
	product := s.sellerProduct.PostProduct(ctx,tx,idSeller,request)

	wg.Add(len(images))
	for _, v := range images {
		go s.sellerProduct.PostProductImages(ctx,tx,product.Id,v,&wg)
	}
	wg.Wait()

	return domain.Products{}
}

func (s *sellerProductServiceImpl) DeleteProduct(ctx context.Context, idProduct int) bool {
	tx,err := s.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	exist := s.sellerProduct.FindById(ctx,tx,idProduct)
	exception.PanicNotFound(exist.Id)
	result := s.sellerProduct.DeleteProduct(ctx,tx,exist.Id)
	return result
}
