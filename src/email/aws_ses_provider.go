package email

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SESService struct {
	Region    string
	FromEmail string
}

func (s *SESService) Send(to, subject, body string) error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(s.Region)})
	if err != nil {
		return err
	}

	svc := ses.New(sess)
	input := &ses.SendEmailInput{
		Source: aws.String(s.FromEmail),
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(to)},
		},
		Message: &ses.Message{
			Subject: &ses.Content{Data: aws.String(subject)},
			Body: &ses.Body{
				Text: &ses.Content{Data: aws.String(body)},
			},
		},
	}

	_, err = svc.SendEmail(input)
	return err
}
