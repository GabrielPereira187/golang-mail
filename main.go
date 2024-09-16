package main

import (
	"bufio"
	"fmt"
	"github/GabrielPereira187/golang-email/entities"
	"github/GabrielPereira187/golang-email/initializers"
	"github/GabrielPereira187/golang-email/sender"
	"os"
	"strconv"
)

func init() {
	initializers.LoadEnvInitializers()
}

func main() {
	if _, err := sender.SendEmail(buildEmail()); err != nil {
		panic(err)
	}

	fmt.Println("Sucesso")

}

func buildEmail() entities.Email {
	var inputEmail string
	fmt.Println("Digite o remetente:")
	_, err := fmt.Scanln(&inputEmail)
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return entities.Email{}
	}

	reader := bufio.NewReader(os.Stdin)

    fmt.Println("Digite o assunto:")  
    inputSubject, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Erro ao ler o assunto:", err)
        return entities.Email{}
    }

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	emailSender := os.Getenv("EMAIL_SENDER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailBody := os.Getenv("EMAIL_BODY")
	resumePath := os.Getenv("RESUME_PATH")


	return entities.Email{
		Host: smtpHost,
		Port: smtpPort,
		Sender: emailSender,
		To: inputEmail,
		Password: emailPassword,
		ResumePath: resumePath,
		Body: emailBody,
		Subject: inputSubject,
	}
}