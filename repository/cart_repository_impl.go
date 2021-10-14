package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/model/web"
)

type cartRepository struct{}

func NewCartRepository() CartRepository {
	return cartRepository{}
}

func (c cartRepository) UpdateItem(ctx context.Context, tx *sql.Tx, request web.CartRequest, idCart int) bool {
	sql := "UPDATE carts SET jumlah_barang = ? WHERE id = ?"
	result,err := tx.ExecContext(ctx,sql,request.Jumlah,idCart)
	exception.PanicIfInternalServerError(err)
	affected,err := result.RowsAffected()
	exception.PanicIfInternalServerError(err)
	return affected > 0
}

func (c cartRepository) PostItem(ctx context.Context, tx *sql.Tx, request web.CartRequest) bool {
	sql := "INSERT INTO carts(id_customer,id_product,jumlah_barang) VALUES(?,?,?)"
	result,err := tx.ExecContext(ctx,sql,request.IdCustomer,request.IdProduct,request.Jumlah)
	exception.PanicIfInternalServerError(err)
	affected,err := result.RowsAffected()
	exception.PanicIfInternalServerError(err)
	return affected > 0
}

func (c cartRepository) GetItems(ctx context.Context, tx *sql.Tx, idCustomer int) []domain.CartProduct {
	sql := "SELECT carts.id AS id_cart, products.id AS id_product , products.nama_barang,products.harga_barang, products.harga_barang * carts.jumlah_barang AS total, (SELECT url FROM product_images WHERE product_images.id_product = products.id LIMIT 1) AS gambar, DATE_FORMAT(carts.created_at,'%w %M %Y') AS tgl FROM carts INNER JOIN products ON products.id  = carts.id_product WHERE carts.id_customer = ? ORDER BY carts.created_at DESC"
	rows,err := tx.QueryContext(ctx,sql,idCustomer)
	exception.PanicIfInternalServerError(err)
	defer rows.Close()

	var items []domain.CartProduct
	for rows.Next() {
		item := domain.CartProduct{}
		err := rows.Scan(&item.IdCart,&item.IdProduct,&item.NamaBarang,&item.HargaBarang,&item.Total,&item.Gambar,&item.CreatedAt)
		exception.PanicIfInternalServerError(err)
		items = append(items, item)
	}
	return items
}

func (c cartRepository) GetCart(ctx context.Context, tx *sql.Tx, idCart int) domain.CartProduct {
	sql := "SELECT carts.id AS id_cart, products.id AS id_product , products.nama_barang,products.harga_barang, products.harga_barang * carts.jumlah_barang AS total, (SELECT url FROM product_images WHERE product_images.id_product = products.id LIMIT 1) AS gambar, DATE_FORMAT(carts.created_at,'%w %M %Y') AS tgl FROM carts INNER JOIN products ON products.id  = carts.id_product WHERE carts.id = ?  ORDER BY carts.created_at DESC"
	rows,err := tx.QueryContext(ctx,sql,idCart)
	exception.PanicIfInternalServerError(err)
	defer rows.Close()

	var item domain.CartProduct
	if rows.Next() {
		item = domain.CartProduct{}
		err := rows.Scan(&item.IdCart,&item.IdProduct,&item.NamaBarang,&item.HargaBarang,&item.Total,&item.Gambar,&item.CreatedAt)
		exception.PanicIfInternalServerError(err)
	}
	return item
}

func (c cartRepository) GetItemsByIdCustomerandProduct(ctx context.Context, tx *sql.Tx, idCustomer int, idProduct int) domain.CartProduct {
	sql := "SELECT carts.id AS id_cart, products.id AS id_product , products.nama_barang,products.harga_barang, products.harga_barang * carts.jumlah_barang AS total, (SELECT url FROM product_images WHERE product_images.id_product = products.id LIMIT 1) AS gambar, DATE_FORMAT(carts.created_at,'%w %M %Y') AS tgl FROM carts INNER JOIN products ON products.id  = carts.id_product WHERE carts.id_customer = ? AND carts.id_product = ? ORDER BY carts.created_at DESC"
	rows,err := tx.QueryContext(ctx,sql,idCustomer,idProduct)
	exception.PanicIfInternalServerError(err)
	defer rows.Close()

	var item domain.CartProduct
	if rows.Next() {
		item = domain.CartProduct{}
		err := rows.Scan(&item.IdCart,&item.IdProduct,&item.NamaBarang,&item.HargaBarang,&item.Total,&item.Gambar,&item.CreatedAt)
		exception.PanicIfInternalServerError(err)
	}
	return item
}

func (c cartRepository) DeleteCart(ctx context.Context, tx *sql.Tx, idCart int) bool {
	sql := "DELETE FROM carts WHERE id = ?"
	result,err := tx.ExecContext(ctx,sql,idCart)
	exception.PanicIfInternalServerError(err)
	affected,err := result.RowsAffected()
	exception.PanicIfInternalServerError(err)
	return affected > 0
}

