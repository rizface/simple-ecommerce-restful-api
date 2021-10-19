package web

type Order struct {
	Invoice string `json:"invoice"`
	IdSeller int `json:"id_seller"`
	IdCustomer int `json:"id_customer"`
	IdProduct int `json:"id_product"`
	Jumlah int `json:"jumlah"`
	Total int `json:"total"`
	Alamat string `json:"alamat"`
}

type OrderRequest struct {
	Items  []OrderProduct `validate:"required" json:"items"`
	Alamat string  `validate:"required" json:"alamat"`
}

type OrderProduct struct {
	IdProduct int `validate:"required" json:"id_product"`
	Jumlah int `validate:"required" json:"jumlah"`
}
