package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"os"
)

type Email struct {
	Subject  string
	Body     string
	Receiver string
}

func SendEmail(email Email) error {
	err := godotenv.Load()
	password := os.Getenv("EMAIL_PASSWORD")
	if err != nil {
		fmt.Println(err)
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "hoanglh1311@yandex.com")
	m.SetHeader("To", email.Receiver)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Body)
	d := gomail.NewDialer("smtp.yandex.com", 465, "hoanglh1311@yandex.com", password)
	if err = d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
