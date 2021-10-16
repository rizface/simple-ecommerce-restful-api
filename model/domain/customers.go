package domain

type Customers struct {
	Id           int    `json:"id,omitempty"`
	NamaCustomer string `validate:"required" json:"nama_customer,omitempty"`
	Email        string `validate:"required" json:"email,omitempty"`
	NoHp         string `validate:"required" json:"no_hp,omitempty"`
	Password     string `validate:"required" json:"password,omitempty"`
	CreatedAt    string `validate:"required" json:"created_at,omitempty"`
	Confirmed    int    `json:"confirmed"`
}
