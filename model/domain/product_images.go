package domain

type ProductImages struct {
	Id int `json:"id,omitempty"`
	IdProduct int `json:"id_product,omitempty"`
	ImageUrl string `json:"url,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}