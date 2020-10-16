package usecase

import (
	"bytes"
	"net/mail"
	"os"
	"untitled/internal/domain"
	"untitled/internal/infrastructure"
)

type SMTPService interface {
	Send(message domain.MailMessage) error
}

type MailTemplateProcessor interface {
	ProcessMailTemplate(template domain.MailTemplate) error
}

func NewMailService(isProductionEnv bool) *MailService {
	ms := new(MailService)
	ms.MailTemplate = &infrastructure.MailTemplateHandler{}
	if isProductionEnv {
		ms.SMTP = &infrastructure.SMTPConnector{
			Host: os.Getenv("SMTP_HOST"),
			Port: os.Getenv("SMTP_PORT"),
			From: os.Getenv("SMTP_FROM"),
			User: os.Getenv("SMTP_USER"),
			Pass: os.Getenv("SMTP_PASS"),
		}
	} else {
		ms.SMTP = &infrastructure.SMTPConnector{
			Host: os.Getenv("SMTP_TEST_HOST"),
			Port: os.Getenv("SMTP_TEST_PORT"),
			From: os.Getenv("SMTP_TEST_FROM"),
			User: os.Getenv("SMTP_TEST_USER"),
			Pass: os.Getenv("SMTP_TEST_PASS"),
		}
	}

	return ms
}

type MailService struct {
	SMTP         SMTPService
	MailTemplate MailTemplateProcessor
}

func (service *MailService) SendRegistrationMail(aggregate domain.AccountAggregate) error {

	type variables struct {
		FullName string
		Subject  string
	}

	subject := "Welcome to paradise"

	mt := domain.MailTemplate{
		Role: "registration",
		Variables: variables{
			FullName: aggregate.Profile.FullName.FullName(),
			Subject:  subject,
		},
		Content: &bytes.Buffer{},
	}

	err := service.MailTemplate.ProcessMailTemplate(mt)
	if err != nil {
		return err
	}

	mm := domain.MailMessage{
		Subject: subject,
	}

	mm.To = append(mm.To, mail.Address{
		Name:    aggregate.Profile.FullName.FullName(),
		Address: aggregate.Identity.Email.Value,
	})

	mm.Content = mt.Content.Bytes()

	err = service.SMTP.Send(mm)
	if err != nil {
		return err
	}

	return nil
}
