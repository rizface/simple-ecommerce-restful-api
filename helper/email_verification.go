package helper

import (
	"gopkg.in/gomail.v2"
	"os"
	"simple-ecommerce-rest-api/app/exception"
)

func SendEmail(url string, to string) {
	LoadConfig()
	url = os.Getenv("HOST") + url
	m := gomail.NewMessage()
	m.SetHeader("From", "malfarizzi13@gmail.com")
	m.SetHeader("To", "malfarizzi33@gmail.com")
	m.SetHeader("Subject", "Verification Email")
	m.SetBody("text/html", "<p>copy and paste link below to verification your account</p><br/>"+url)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	err := dialer.DialAndSend(m)
	exception.PanicIfInternalServerError(err)
}
