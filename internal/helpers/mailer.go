package helpers

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(toEmail, subject, body string) error {

	//Domain e-mail
	from := mail.NewEmail("Taskly", "me@pedrooyarzun.xyz")
	to := mail.NewEmail(toEmail, toEmail)

	message := mail.NewSingleEmail(from, subject, to, body, body)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	_, err := client.Send(message)

	return err
}
