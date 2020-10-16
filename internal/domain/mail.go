package domain

import (
	"bytes"
	"net/mail"
)

type MailMessage struct {
	To      []mail.Address
	Content []byte
	Subject string
}

type MailTemplate struct {
	Role      string
	Variables interface{}
	Content   *bytes.Buffer
}

type MailService interface {
	SendRegistrationMail(aggregate AccountAggregate) error
}
