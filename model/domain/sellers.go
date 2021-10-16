package domain

type Seller struct {
	Id         int    `json:"id"`
	NamaToko   string `json:"nama_toko,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	AlamatToko string `json:"alamat_toko,omitempty"`
	Deskripsi  string `json:"deskripsi,omitempty"`
	CreatedAt  string `json:"omitempty,created_at"`
	Confirmed  int    `json:"confirmed"`
}
