package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"os"
)

func SendEmail(receiver string, subject string, body string) error {
	err := godotenv.Load()
	password := os.Getenv("EMAIL_PASSWORD")
	if err != nil {
		fmt.Println(err)
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "hoanglh1311@yandex.com")
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer("smtp.yandex.com", 465, "hoanglh1311@yandex.com", password)
	if err = d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
