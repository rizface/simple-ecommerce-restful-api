package web

type RequestCustomer struct{
	NamaCustomer string `validate:"required" json:"nama_customer,omitempty"`
	Email string `validate:"required" json:"email,omitempty"`
	NoHp string `validate:"required" json:"no_hp,omitempty"`
	Password string `validate:"required" json:"password,omitempty"`
}
