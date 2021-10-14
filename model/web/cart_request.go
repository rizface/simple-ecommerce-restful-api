package web

type CartRequest struct {
	IdCustomer int `validate:"required" json:"id_customer"`
	IdProduct  int `validate:"required" json:"id_product"`
	Jumlah     int `validate:"required" json:"jumlah"`
}
