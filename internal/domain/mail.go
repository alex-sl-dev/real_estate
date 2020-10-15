package domain

import "bytes"

type MailMessage struct {
	To []string
	Content []byte
	Subject string
}

type MailTemplate struct {
	Role string
	Variables interface{}
	Content *bytes.Buffer
}


type MailService interface {
	SendRegistrationMail(aggregate AccountAggregate) error
}