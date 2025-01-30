package service

import "github.com/sahidhossen/synmail/src/config"

type EmailClient struct {
	Gateway   MailInterface
	fromEmail string
	fromName  string
}

type MailInterface interface{}

func NewEmailClient(cfg *config.Config, emailGateway MailInterface) *EmailClient {
	return &EmailClient{
		Gateway:   emailGateway,
		fromEmail: cfg.MailFrom,
		fromName:  cfg.MailFromName,
	}
}
