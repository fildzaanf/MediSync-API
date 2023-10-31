package helper

import (
	"app/model/domain"
	"fmt"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

type Mail struct {
	smtpHost     string
	smtpPort     string
	smtpEmail    string
	smtpPassword string
}

func SendEmailController(email, subject string, medicine domain.Medicine) error {

	mail := Mail{
		smtpHost:     os.Getenv("SMTP_HOST"),
		smtpPort:     os.Getenv("SMTP_PORT"),
		smtpEmail:    os.Getenv("SMTP_EMAIL"),
		smtpPassword: os.Getenv("SMTP_PASSWORD"),
	}

	fmt.Println(email, subject, medicine)
	fmt.Println(mail)
	name := medicine.Name
	amount := medicine.Amount
	details := medicine.Details

	
	body := fmt.Sprintf("<h3>[MediSync]Reminder<h3>\n<h4>Medicine 	:	%s<h4>\n <h4>Amounts		: 	%d<h4>\n  <h4>Details 		:	%s<h4>\n <h3>Get Well Soon!<h3>", name, amount, details)

	port, err := strconv.Atoi(mail.smtpPort)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mail.smtpEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		mail.smtpHost,
		port,
		mail.smtpEmail,
		mail.smtpPassword,
	)

	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
