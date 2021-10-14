package repository

import (
	"context"
	"database/sql"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/model/domain"
)

type productImagesRepositoryImpl struct{}

func NewProductImagesRepoImpl() ProductImagesRepository {
	return productImagesRepositoryImpl{}
}

func (p productImagesRepositoryImpl) GetByProductId(ctx context.Context, tx *sql.Tx, idProduct int) []domain.ProductImages {
	sql := "SELECT url FROM product_images WHERE id_product = ?"
	rows, err := tx.QueryContext(ctx, sql, idProduct)
	exception.PanicIfInternalServerError(err)
	images := []domain.ProductImages{}
	defer rows.Close()
	for rows.Next() {
		image := domain.ProductImages{}
		err := rows.Scan(&image.ImageUrl)
		exception.PanicIfInternalServerError(err)
		images = append(images, image)
	}
	return images
}
