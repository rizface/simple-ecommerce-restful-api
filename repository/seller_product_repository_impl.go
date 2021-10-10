package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
	"sync"
)

type sellerProductRepositoryImpl struct{}

func NewSellerProductRepositoryImpl() SellerProductRepository{
	return sellerProductRepositoryImpl{}
}


func (s sellerProductRepositoryImpl) GetProducts(ctx context.Context, tx *sql.Tx, idSeller int) []domain.Products {
	var products []domain.Products
	sql := "SELECT id,id_seller,nama_barang,harga_barang,stok_barang,deskripsi,DATE_FORMAT(created_at, '%w %M %Y') FROM products WHERE id_seller = ?"
	rows,err := tx.QueryContext(ctx,sql,idSeller)
	exception.PanicIfInternalServerError(err)
	for rows.Next() {
		each := domain.Products{}
		err := rows.Scan(&each.Id,&each.IdSeller,&each.NamaBarang,&each.HargaBarang,&each.StokBarang,&each.Deskripsi,&each.CreatedAt)
		exception.PanicIfInternalServerError(err)
		products = append(products,each)
	}
	return products
}

func (s sellerProductRepositoryImpl) PostProduct(ctx context.Context, tx *sql.Tx, idSeller int, request web.NewProduct) domain.Products {
	sql := "INSERT INTO products(id_seller,nama_barang,harga_barang,stok_barang,deskripsi) VALUES(?,?,?,?,?)"
	result,err := tx.ExecContext(ctx,sql, idSeller,request.NamaBarang,request.HargaBarang,request.Stokbarang,request.Deskripsi)
	exception.PanicIfInternalServerError(err)
	id,_ := result.LastInsertId()
	return domain.Products{
		Id:          int(id),
		IdSeller:    idSeller,
		NamaBarang:  request.NamaBarang,
		HargaBarang: request.HargaBarang,
		StokBarang:  request.Stokbarang,
		Deskripsi:   request.Deskripsi,
	}
}

func (s sellerProductRepositoryImpl) PostProductImages(ctx context.Context, tx *sql.Tx, idProduct int, link string, wg *sync.WaitGroup) {
	defer wg.Done()
	sql := "INSERT INTO product_images(id_product,url) VALUES(?,?)"
	_, err := tx.ExecContext(ctx,sql,idProduct,link)
	exception.PanicIfInternalServerError(err)
}
