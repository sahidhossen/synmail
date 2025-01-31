package email

import (
	"fmt"
	"net/smtp"
)

type SMTPService struct {
	Host      string
	Port      int
	Username  string
	Password  string
	FromEmail string
}

func (s *SMTPService) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	msg := []byte("From: " + s.FromEmail + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n")

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	return smtp.SendMail(addr, auth, s.FromEmail, []string{to}, msg)
}
