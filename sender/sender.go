package sender

import (
	"github/GabrielPereira187/golang-email/entities"
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(email entities.Email) (bool, error) {

	log.Println("Enviando......")


	mail := gomail.NewMessage()
	mail.SetHeader("From", email.Sender)
	mail.SetHeader("To", email.To)
	mail.SetHeader("Subject", email.Subject)
	mail.SetBody("text/plain", email.Body)

	mail.Attach(email.ResumePath)

	dial := gomail.NewDialer(email.Host, email.Port, email.Sender, email.Password)

	if err := dial.DialAndSend(mail); err != nil {
		return false, err
	}

	return true, nil
}