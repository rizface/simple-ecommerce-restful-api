package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type sellerRepositoryImpl struct{}

func NewSellerRepoImpl() SellerRepository {
	return sellerRepositoryImpl{}
}

func (s sellerRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, request web.RequestSeller) int64 {
	sql := "INSERT INTO sellers(nama_toko,email,password,alamat_toko,deskripsi) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, sql, request.NamaToko, request.Email, request.Password, request.AlamatToko, request.Deskripsi)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	return id
}

func (s sellerRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) domain.Seller {
	sql := "SELECT id,nama_toko,email,password,alamat_toko,deskripsi, DATE_FORMAT(created_at, '%w %M %Y') FROM sellers WHERE email = ?"
	rows, err := tx.QueryContext(ctx, sql, email)
	helper.PanicIfError(err)
	defer rows.Close()
	seller := domain.Seller{}
	if rows.Next() {
		err := rows.Scan(&seller.Id, &seller.NamaToko, &seller.Email, &seller.Password, &seller.AlamatToko, &seller.Deskripsi, &seller.CreatedAt)
		helper.PanicIfError(err)
	}
	return seller
}

func (s sellerRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) domain.Seller {
	sql := "SELECT id,nama_toko,email,password,alamat_toko,deskripsi, DATE_FORMAT(created_at, '%w %M %Y') FROM sellers WHERE nama_toko = ?"
	rows, err := tx.QueryContext(ctx, sql, name)
	helper.PanicIfError(err)
	defer rows.Close()
	seller := domain.Seller{}
	if rows.Next() {
		err := rows.Scan(&seller.Id, &seller.NamaToko, &seller.Email, &seller.Password, &seller.AlamatToko, &seller.Deskripsi, &seller.CreatedAt)
		helper.PanicIfError(err)
	}
	return seller
}
