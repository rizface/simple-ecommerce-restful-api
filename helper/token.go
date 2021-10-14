package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/model/domain"
	"time"
)

var SellerSecret = []byte("seller-secret")
var CustomerSecret = []byte("customer-secret")

type SellerCustom struct {
	Id        int    `json:"id"`
	NamaToko  string `json:"nama_toko"`
	Email     string `json:"email"`
	Deskripsi string `json:"deskripsi"`
	Seller    bool   `json:"seller"`
	CreatedAt string `json:"created_at"`
	jwt.RegisteredClaims
}
type CustomerCustom struct {
	Id           int    `json:"id"`
	NamaCustomer string `json:"nama_toko"`
	Email        string `json:"email"`
	NoHp         string `json:"no_hp"`
	CreatedAt    string `json:"created_at"`
	jwt.RegisteredClaims
}

func GenerateTokenCustomer(customer domain.Customers) string {
	customerClaims := CustomerCustom{
		Id:           customer.Id,
		NamaCustomer: customer.NamaCustomer,
		Email:        customer.Email,
		NoHp:         customer.NoHp,
		CreatedAt:    customer.CreatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Muhammad Al Farizzi",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customerClaims)
	strToken, err := token.SignedString(CustomerSecret)
	exception.PanicIfInternalServerError(err)
	return strToken
}

func GenerateTokenSeller(seller domain.Seller) string {
	sellerClaims := SellerCustom{
		Id:        seller.Id,
		NamaToko:  seller.NamaToko,
		Email:     seller.Email,
		Deskripsi: seller.Deskripsi,
		CreatedAt: seller.CreatedAt,
		Seller:    true,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Muhammad Al Farizzi",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, sellerClaims)
	sellerToken, err := token.SignedString(SellerSecret)
	exception.PanicIfInternalServerError(err)
	return sellerToken
}

func VerifyToken(sellerToken string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(sellerToken, &SellerCustom{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if ok {
			return SellerSecret, nil
		} else {
			return nil, errors.New("Signing Methos Is Invalid")
		}
	})
	if err != nil {
		return nil, err
	}
	claims, claimsOK := token.Claims.(*SellerCustom)
	if !claimsOK || token.Valid == false {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}

func VerifyTokenCustomer(customerToken string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(customerToken, &CustomerCustom{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if ok {
			return CustomerSecret, nil
		}
		return nil, errors.New("token format is invalid")
	})
	if err != nil {
		return nil, err
	}

	claims, claimsOK := token.Claims.(*CustomerCustom)
	if claimsOK && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is invalid")
}
