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
)

type sellerServiceImpl struct {
	db *sql.DB
	validate *validator.Validate
	sellerRepo repository.SellerRepository
}

func NewSellerServiceImpl(validate *validator.Validate, db *sql.DB, sellerRepo repository.SellerRepository) SellerService {
	return &sellerServiceImpl{validate: validate,db:db, sellerRepo: sellerRepo}
}

func (s *sellerServiceImpl) Register(ctx context.Context, request web.RequestSeller) domain.Seller {
	err := s.validate.Struct(request)
	exception.PanicBadRequest(err)

	tx,err := s.db.Begin()
	helper.PanicIfError(err)

	existEmail := s.sellerRepo.FindByEmail(ctx,tx,request.Email)
	exception.PanicDuplicate(existEmail.Id, existEmail.Email + " Sudah Digunakan")

	existStore := s.sellerRepo.FindByName(ctx,tx,request.NamaToko)
	exception.PanicDuplicate(existStore.Id, existStore.NamaToko + " Suda Terdaftar")

	defer helper.CommitOrRollback(tx)
	sellerId := s.sellerRepo.Register(ctx,tx,request)
	seller := domain.Seller{
		Id:         int(sellerId),
		NamaToko:   request.NamaToko,
		Email:      request.Email,
		AlamatToko: request.AlamatToko,
		Deskripsi:  request.Deskripsi,
	}
	return seller
}

func (s *sellerServiceImpl) Login(ctx context.Context, request web.RequestSeller) string {
	tx,err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	existEmail := s.sellerRepo.FindByEmail(ctx,tx,request.Email)
	exception.PanicNotFound(existEmail.Id)

	err = helper.Compare(request.Password,existEmail.Password)
	exception.PanicUnauthorized(err)
	token := helper.GenerateTokenSeller(existEmail)
	return token
}

