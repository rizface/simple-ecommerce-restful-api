package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/domain"
	"simple-ecommerce-rest-api/repository"
	"strconv"
	"time"
)

type customerProductServiceImpl struct {
	db                *sql.DB
	productImagesRepo repository.ProductImagesRepository
	productRepo       repository.CustomerProduct
}

func NewCustomerProductServiceImpl(db *sql.DB, productImagesRepo repository.ProductImagesRepository, productRepo repository.CustomerProduct) CustomerProductService {
	return customerProductServiceImpl{
		db:                db,
		productImagesRepo: productImagesRepo,
		productRepo:       productRepo,
	}
}

func (c customerProductServiceImpl) setImages(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	product.Gambar = c.productImagesRepo.GetByProductId(ctx, tx, product.Id)
	return product
}

func (c customerProductServiceImpl) Get(ctx context.Context) []domain.Products {
	tx, err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	products := c.productRepo.Get(ctx, tx)
	for _, v := range products {
		product := c.setImages(ctx, tx, v)
		products = append(products, product)
	}
	products = products[len(products)/2 : len(products)]
	return products
}

func (c customerProductServiceImpl) GetDetail(ctx context.Context, idProduct int) domain.Products {
	tx, err := c.db.Begin()
	exception.PanicIfInternalServerError(err)
	defer helper.CommitOrRollback(tx)
	var product domain.Products
	redisKey := "product" + strconv.Itoa(idProduct)
	productCache, err := helper.Rdb.Get(ctx, redisKey).Result()

	if err == redis.Nil {
		product = c.productRepo.GetDetail(ctx, tx, idProduct)
		exception.PanicNotFound(product.Id)
		product.Gambar = c.productImagesRepo.GetByProductId(ctx, tx, product.Id)
		productJson, _ := json.Marshal(product)
		_, err := helper.Rdb.SetEX(ctx, redisKey, productJson, time.Duration(60)*time.Second).Result()
		exception.PanicIfInternalServerError(err)
		return product
	} else {
		json.Unmarshal([]byte(productCache), &product)
	}
	return product
}
