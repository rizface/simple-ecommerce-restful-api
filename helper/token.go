package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"simple-ecommerce-rest-api/model/domain"
	"time"
)

var SellerSecret = []byte("seller-secret")
type SellerCustom struct{
	Id int `json:"id"`
	NamaToko string `json:"nama_toko"`
	Email string `json:"email"`
	Deskripsi string `json:"deskripsi"`
	CreatedAt string `json:"created_at"`
	jwt.RegisteredClaims
}

func GenerateTokenSeller(seller domain.Seller) string {
	sellerClaims := SellerCustom{
		Id:               seller.Id,
		NamaToko:         seller.NamaToko,
		Email:            seller.Email,
		Deskripsi:        seller.Deskripsi,
		CreatedAt:        seller.CreatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "Muhammad Al Farizzi",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(10) * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,sellerClaims)
	sellerToken,err := token.SignedString(SellerSecret)
	PanicIfError(err)
	return sellerToken
}

//func VerifyToken() {
//
//}
