package domain

type Products struct {
	Id int `json: "id"`
	IdSeller int `json:"id_seller"`
	NamaBarang string `json:"nama_barang"`
	HargaBarang int `json:"harga_barang"`
	StokBarang int `json:"stok_barang"`
	Deskripsi string `json:"deskripsi"`
	CreatedAt string `json:"created_at"`
}