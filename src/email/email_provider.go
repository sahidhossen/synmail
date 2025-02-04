package email

import (
	"errors"

	"github.com/sahidhossen/synmail/src/config"
)

type ProviderType string

const (
	SMTP     ProviderType = "smtp"
	SendGrid ProviderType = "sendgrid"
	Mailgun  ProviderType = "mailgun"
	SES      ProviderType = "ses"
)

func NewEmailService(provider ProviderType, cfg *config.Config) (EmailService, error) {
	switch provider {
	case SMTP:
		return &SMTPService{
			Host:      cfg.MailHost,
			Port:      cfg.MailPort,
			Username:  cfg.MailUsername,
			Password:  cfg.MailPass,
			FromEmail: cfg.MailFrom,
		}, nil
	case SendGrid:
		return &SendGridService{APIKey: cfg.MailAPIKey, FromEmail: cfg.MailFrom}, nil
	case Mailgun:
		return &MailgunService{Domain: cfg.MailGunDomain, APIKey: cfg.MailAPIKey, FromEmail: cfg.MailFrom}, nil
	case SES:
		return &SESService{Region: cfg.AWSRegion}, nil
	default:
		return nil, errors.New("unsupported email provider")
	}
}
