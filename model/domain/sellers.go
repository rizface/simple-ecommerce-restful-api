package domain

type Seller struct {
	Id int `json:"id"`
	NamaToko string `json:"omitempty, nama_toko"`
	Email string `json:"omitempty, email"`
	Password string `json:"omitempty, password"`
	AlamatToko string `json:"omitempty, alamat_toko"`
	Deskripsi string `json:"omitempty, deskripsi"`
	CreatedAt string `json:"omitempty,created_at"`
}
