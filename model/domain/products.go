package domain

type Products struct {
	Id          int             `json: "id"`
	IdSeller    int             `json:"id_seller,omitempty"`
	NamaBarang  string          `json:"nama_barang,omitempty"`
	HargaBarang int             `json:"harga_barang,omitempty"`
	StokBarang  int             `json:"stok_barang,omitempty"`
	Deskripsi   string          `json:"deskripsi,omitempty"`
	Gambar      []ProductImages `json:"gambar,omitempty"`
	CreatedAt   string          `json:"created_at,omitempty"`
}
