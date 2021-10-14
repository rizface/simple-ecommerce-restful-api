package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type customerRepositoryImpl struct{}

func NewCustomerRepositoryImpl() CustomerRepository {
	return customerRepositoryImpl{}
}

func (c customerRepositoryImpl) RegisterCustomer(ctx context.Context, tx *sql.Tx, request web.RequestCustomer) bool {
	sql := "INSERT INTO customers(nama_customer,email,no_hp,password) VALUES(?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, request.NamaCustomer, request.Email, request.NoHp, request.Password)
	exception.PanicIfInternalServerError(err)
	affected, err := result.RowsAffected()
	exception.PanicIfInternalServerError(err)
	return affected > 0
}

func (c customerRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) domain.Customers {
	sql := "SELECT id,nama_customer,email,no_hp,password,DATE_FORMAT(created_at,'%w %M %Y') FROM customers WHERE email = ?"
	row, err := tx.QueryContext(ctx, sql, email)
	exception.PanicIfInternalServerError(err)
	defer row.Close()
	customer := domain.Customers{}
	if row.Next() {
		err := row.Scan(&customer.Id, &customer.NamaCustomer, &customer.Email, &customer.NoHp, &customer.Password, &customer.CreatedAt)
		exception.PanicIfInternalServerError(err)
	}

	return customer
}
