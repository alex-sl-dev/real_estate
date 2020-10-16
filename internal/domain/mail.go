package domain

import (
	"bytes"
	"net/mail"
)

const (
	MailRegistrationTpl        = "registration"
	MailAddressVerificationTpl = "mailAddressValidation"
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
	SendConfirmAddressMail(emailAddress, code string) error
}
