package sender

import "github.com/aws/aws-sdk-go/service/ses"

// Service interface for the sender service
type Service interface {
	Build(to, html, subject string) *ses.SendEmailInput
	Send(input *ses.SendEmailInput) (*ses.SendEmailOutput, error)
}
