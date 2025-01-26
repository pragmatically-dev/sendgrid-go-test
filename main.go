package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	from := mail.NewEmail("NR Entropía", "entropia.talleres.dev@gmail.com")
	subject := "Este es un test de envio"
	to := mail.NewEmail("Santiago", "pragmatically.dev@gmail.com")
	plainTextContent := "Este es un test de envio"
	htmlContent := "<strong>Test de envio desde golang</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)

	}
}
