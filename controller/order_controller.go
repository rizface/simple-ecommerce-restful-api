package controller

import "net/http"

type OrderController interface {
	GetOrders(w http.ResponseWriter,r *http.Request)
	PostOrders(w http.ResponseWriter, r *http.Request)
}
