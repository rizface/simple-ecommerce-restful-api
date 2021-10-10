package web

type NewProduct struct {
	NamaBarang string `validate:"required" json:"nama_barang"`
	HargaBarang int `validate:"required" json:"harga_barang"`
	Stokbarang int `validate:"required" json:"stok_barang"`
	Deskripsi string `validate:"required" json:"deskripsi"`
	Gambar []string `json:"gambar"`
}
