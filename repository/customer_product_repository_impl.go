package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/model/domain"
)

type customerProductRepoImpl struct{}

func NewCustomerProductRepoImpl() CustomerProduct {
	return customerProductRepoImpl{}
}

func (c customerProductRepoImpl) Get(ctx context.Context, tx *sql.Tx) []domain.Products {
	sql := "SELECT id,nama_barang,harga_barang,stok_barang,deskripsi FROM products"
	rows, err := tx.QueryContext(ctx, sql)
	exception.PanicIfInternalServerError(err)
	defer rows.Close()
	products := []domain.Products{}
	for rows.Next() {
		product := domain.Products{}
		err := rows.Scan(&product.Id, &product.NamaBarang, &product.HargaBarang, &product.StokBarang, &product.Deskripsi)
		exception.PanicIfInternalServerError(err)
		products = append(products, product)
	}
	return products
}

func (c customerProductRepoImpl) GetDetail(ctx context.Context, tx *sql.Tx, idProduct int) domain.Products {
	sql := "SELECT id,nama_barang,harga_barang,stok_barang,deskripsi FROM products WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, idProduct)
	exception.PanicIfInternalServerError(err)
	defer rows.Close()
	product := domain.Products{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.NamaBarang, &product.HargaBarang, &product.StokBarang, &product.Deskripsi)
		exception.PanicIfInternalServerError(err)
	}
	return product
}
