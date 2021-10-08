package controller

import "net/http"

type SellerController interface {
	Register(w http.ResponseWriter, r *http.Request)
}
