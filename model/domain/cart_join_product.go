package domain

type CartProduct struct {
	IdCart int `json:"id_cart"`
	IdProduct int `json:"id_product"`
	NamaBarang string `json:"nama_barang"`
	HargaBarang int `json:"harga_barang"`
	Total int `json:"total"`
	Gambar string `json:"gambar"`
	CreatedAt string `json:"created_at"`
}
