package app

import (
	"github.com/gorilla/mux"
)

var Mux *mux.Router = mux.NewRouter()

const (
	SELLER_REGITER = "/seller/register"
	SELLER_LOGIN = "/seller/login"
	SELLER_PRODUCT = "/seller/products"
	SELLER_PROUDUCT_MANIPULATION = "/seller/products/{idProduct}"
)
