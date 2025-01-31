package email

type EmailService interface {
	Send(to, subject, body string) error
}
