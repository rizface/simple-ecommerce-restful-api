package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/repository"
)

type customerServiceImpl struct {
	db *sql.DB
	validate *validator.Validate
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
	tx,err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	exist := c.repository.FindByEmail(ctx,tx,request.Email)
	exception.PanicDuplicate(exist.Id,"email sudah digunakan")
	request.Password = helper.Hash(request.Password)
	result := c.repository.RegisterCustomer(ctx,tx,request)
	return result
}

func (c customerServiceImpl) LoginCustomer(ctx context.Context, request web.RequestCustomer) string {
	tx,err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	exist := c.repository.FindByEmail(ctx,tx,request.Email)
	exception.PanicNotFound(exist.Id)
	err = helper.Compare(request.Password,exist.Password)
	exception.PanicUnauthorized(err)
	token := helper.GenerateTokenCustomer(exist)
	return token
}



