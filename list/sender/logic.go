package sender

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type service struct {
	client *ses.SES
	from   string
}

// New returns a new service
func New(region, from string) (Service, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			// Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
		},
	)
	if err != nil {
		return &service{}, err
	}

	return &service{ses.New(sess), from}, nil
}

func (s *service) Build(to, html, subject string) *ses.SendEmailInput {
	charSer := "UTF-8"

	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSer),
					Data:    aws.String(html),
				},
				// Text: &ses.Content{
				// 	Charset: aws.String(charSer),
				// 	Data:    aws.String(TextBody),
				// },
			},
			Subject: &ses.Content{
				Charset: aws.String(charSer),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(s.from),
	}
}

func (s *service) Send(input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	return s.client.SendEmail(input)
}
