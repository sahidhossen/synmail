package email

import (
	"context"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunService struct {
	FromEmail string
	Domain    string
	APIKey    string
}

func (m *MailgunService) Send(to, subject, body string) error {
	mg := mailgun.NewMailgun(m.Domain, m.APIKey)
	message := mailgun.NewMessage(m.FromEmail, subject, body, to)
	_, _, err := mg.Send(context.Background(), message)
	return err
}
