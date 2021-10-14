package controller

import "net/http"

type CustomerProductController interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
}
