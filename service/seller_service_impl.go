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

type sellerServiceImpl struct {
	db         *sql.DB
	validate   *validator.Validate
	sellerRepo repository.SellerRepository
}

func NewSellerServiceImpl(validate *validator.Validate, db *sql.DB, sellerRepo repository.SellerRepository) SellerService {
	return &sellerServiceImpl{validate: validate, db: db, sellerRepo: sellerRepo}
}

func (s *sellerServiceImpl) Register(ctx context.Context, request web.RequestSeller) domain.Seller {
	err := s.validate.Struct(request)
	exception.PanicBadRequest(err)

	tx, err := s.db.Begin()
	helper.PanicIfError(err)

	existEmail := s.sellerRepo.FindByEmail(ctx, tx, request.Email)
	exception.PanicDuplicate(existEmail.Id, existEmail.Email+" Sudah Digunakan")

	existStore := s.sellerRepo.FindByName(ctx, tx, request.NamaToko)
	exception.PanicDuplicate(existStore.Id, existStore.NamaToko+" Sudah Terdaftar")

	defer helper.CommitOrRollback(tx)
	sellerId := s.sellerRepo.Register(ctx, tx, request)
	seller := domain.Seller{
		Id:         int(sellerId),
		NamaToko:   request.NamaToko,
		Email:      request.Email,
		AlamatToko: request.AlamatToko,
		Deskripsi:  request.Deskripsi,
	}

	// send verification email
	token := helper.GenerateTokenSeller(seller)
	helper.SendEmail(":8080/seller/"+token, seller.Email)
	// send verification email
	return seller
}

func (s *sellerServiceImpl) Login(ctx context.Context, request web.RequestSeller) string {
	tx, err := s.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	existEmail := s.sellerRepo.FindByEmail(ctx, tx, request.Email)
	exception.PanicNotFound(existEmail.Id)

	err = helper.Compare(request.Password, existEmail.Password)
	exception.PanicUnauthorized(err)
	token := helper.GenerateTokenSeller(existEmail)
	return token
}

func (s *sellerServiceImpl) Confirm(ctx context.Context, token string) string {
	claims, err := helper.VerifyToken(token)
	seller, sellerOK := claims.(*helper.SellerCustom)
	if err != nil || !sellerOK {
		exception.PanicBadRequest(errors.New("token is invalid"))
	}
	tx, err := s.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	existSeller := s.sellerRepo.FindByEmail(ctx, tx, seller.Email)
	exception.PanicNotFound(existSeller.Id)
	success := s.sellerRepo.Confirm(ctx, tx, existSeller.Id)
	if success {
		return "seller account confirmation is success"
	}
	return "seller account confirmation is failed"
}
