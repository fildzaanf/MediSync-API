package helper

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

type Mail struct {
	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_EMAIL    string
	SMTP_PASSWORD string
}

func SendEmailController(email, subject, body string) error {

	mail := Mail{
		SMTP_HOST:     os.Getenv("SMTP_HOST"),
		SMTP_PORT:     os.Getenv("SMTP_PORT"),
		SMTP_EMAIL:    os.Getenv("SMTP_EMAIL"),
		SMTP_PASSWORD: os.Getenv("SMTP_PASSWORD"),
	}

    fmt.Println(email, subject, body)
    fmt.Println(mail)

    port, err := strconv.Atoi(mail.SMTP_PORT)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mail.SMTP_EMAIL)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

    dialer := gomail.NewDialer(
		mail.SMTP_HOST,
		port,
		mail.SMTP_EMAIL,
		mail.SMTP_PASSWORD,
	)

	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
