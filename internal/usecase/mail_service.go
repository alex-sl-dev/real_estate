package usecase

import (
	"bytes"
	"untitled/internal/domain"
)

type SMTPService interface {
	Send(message domain.MailMessage) error
}

type MailTemplateProcessor interface {
	ProcessMailTemplate(template domain.MailTemplate) error
}

type MailService struct {
	SMTP         SMTPService
	MailTemplate MailTemplateProcessor
}

func (service *MailService) SendRegistrationMail(aggregate domain.AccountAggregate) error {

	type variables struct{
		FullName string
	}

	mt := domain.MailTemplate{
		//Subject: "Welcome to paradise",
		Role: "registration",
		Variables: variables{
			FullName: aggregate.Profile.FullName.FullName(),
		},
		Content: &bytes.Buffer{},
	}

	err := service.MailTemplate.ProcessMailTemplate(mt)
	if err != nil {
		return err
	}

	mm := domain.MailMessage{}
	mm.Content = mt.Content.Bytes()

	err = service.SMTP.Send(mm)
	if err != nil {
		return err
	}

	return nil
}
