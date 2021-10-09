package web

type RequestSeller struct {
	NamaToko string `validate:"required" json:"nama_toko"`
	Email string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	AlamatToko string `validate:"required" json:"alamat_toko"`
	Deskripsi string `validate:"required" json:"deskripsi"`
}
