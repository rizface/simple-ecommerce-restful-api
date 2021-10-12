package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"simple-ecommerce-rest-api/model/domain"
	"time"
)

// Seller Secret
var SellerSecret = []byte("seller-secret")
// Customer Secret
var CustomerSecret = []byte("customer-secret")

// Seller Custom Claims
type SellerCustom struct{
	Id			int `json:"id"`
	NamaToko 	string `json:"nama_toko"`
	Email 		string `json:"email"`
	Deskripsi	string `json:"deskripsi"`
	Seller 		bool `json:"seller"`
	CreatedAt 	string `json:"created_at"`
	jwt.RegisteredClaims
}
// Custoemr Customer Claims
type CustomerCustom struct{}

// Seller Token Generator
func GenerateTokenSeller(seller domain.Seller) string {
	sellerClaims := SellerCustom{
		Id:               seller.Id,
		NamaToko:         seller.NamaToko,
		Email:            seller.Email,
		Deskripsi:        seller.Deskripsi,
		CreatedAt:        seller.CreatedAt,
		Seller: true,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "Muhammad Al Farizzi",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,sellerClaims)
	sellerToken,err := token.SignedString(SellerSecret)
	PanicIfError(err)
	return sellerToken
}

func VerifyToken(sellerToken string) (interface{},error) {
	token,err := jwt.ParseWithClaims(sellerToken,&SellerCustom{}, func(token *jwt.Token) (interface{}, error) {
		_,ok := token.Method.(*jwt.SigningMethodHMAC)
		if ok {
			return SellerSecret,nil
		} else {
			return nil, errors.New("Signing Methos Is Invalid")
		}
	})
	if err != nil {
		return nil,err
	}
	claims,claimsOK := token.Claims.(*SellerCustom)
	if !claimsOK || token.Valid == false {
		return  nil,errors.New("token is invalid")
	}
	return claims,nil
}

// Customer Token Validator