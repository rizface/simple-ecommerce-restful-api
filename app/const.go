package app

import (
	"github.com/gorilla/mux"
)

var Mux *mux.Router = mux.NewRouter()

const (
	// Customer
	CUSTOMER_REGISTER = "/customer/register"
	CUSTOMER_LOGIN = "/customer/login"
	PRODUCTS = "/products"
	PRODUCT_DETAIL = "/products/{idProduct}"

	// Seller
	SELLER_REGITER = "/seller/register"
	SELLER_LOGIN = "/seller/login"
	SELLER_PRODUCT = "/seller/products"
	SELLER_PROUDUCT_MANIPULATION = "/seller/products/{idProduct}"
	UPLOAD_IMAGES = "http://localhost:8081/api/save"
)
