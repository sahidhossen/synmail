package email

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridService struct {
	FromEmail string
	APIKey    string
}

func (s *SendGridService) Send(to, subject, body string) error {
	from := mail.NewEmail("SynMail", s.FromEmail)
	toEmail := mail.NewEmail("", to)
	message := mail.NewSingleEmail(from, subject, toEmail, body, body)
	client := sendgrid.NewSendClient(s.APIKey)
	_, err := client.Send(message)
	return err
}
