package usecase

import (
	"os"
	"untitled/internal/infrastructure"
)

func NewMailService(isProductionEnv bool) *MailService {

	ms := new(MailService)

	ms.MailTemplate = &infrastructure.MailTemplateHandler{

	}

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
