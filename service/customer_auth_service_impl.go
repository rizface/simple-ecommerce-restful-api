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

type customerServiceImpl struct {
	db         *sql.DB
	validate   *validator.Validate
	repository repository.CustomerRepository
}

func NewCustomerServiceImpl(db *sql.DB, validate *validator.Validate, repository repository.CustomerRepository) CustomerAuthService {
	return customerServiceImpl{
		db:         db,
		validate:   validate,
		repository: repository,
	}
}

func (c customerServiceImpl) RegisterCustomer(ctx context.Context, request web.RequestCustomer) bool {
	err := c.validate.Struct(request)
	exception.PanicBadRequest(err)
	tx, err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	exist := c.repository.FindByEmail(ctx, tx, request.Email)
	exception.PanicDuplicate(exist.Id, "email sudah digunakan")
	request.Password = helper.Hash(request.Password)
	result := c.repository.RegisterCustomer(ctx, tx, request)

	// send email verification
	token := helper.GenerateTokenCustomer(domain.Customers{
		Id:           int(result),
		NamaCustomer: request.NamaCustomer,
		Email:        request.Email,
		NoHp:         request.NoHp,
		Confirmed:    0,
	})
	helper.SendEmail(":8080/customer/"+token, request.Email)
	// send email verification

	return result > 0
}

func (c customerServiceImpl) LoginCustomer(ctx context.Context, request web.RequestCustomer) string {
	tx, err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	exist := c.repository.FindByEmail(ctx, tx, request.Email)
	exception.PanicNotFound(exist.Id)
	err = helper.Compare(request.Password, exist.Password)
	exception.PanicUnauthorized(err)
	token := helper.GenerateTokenCustomer(exist)
	return token
}

func (c customerServiceImpl) Confirm(ctx context.Context, token string) string {
	claims,err := helper.VerifyTokenCustomer(token)
	customer,customerOK := claims.(*helper.CustomerCustom)

	if err != nil || !customerOK {
		exception.PanicBadRequest(errors.New("token invalid"))
	}

	tx,err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	existCustomer := c.repository.FindByEmail(ctx,tx,customer.Email)
	exception.PanicNotFound(existCustomer.Id)
	success := c.repository.UpdateConfirmed(ctx,tx,existCustomer.Id)
	if success {
		return "account confirmation success"
	}
	return "account confirmation failed"
}
