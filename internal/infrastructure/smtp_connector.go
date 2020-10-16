package infrastructure

import (
	"bytes"
	"log"
	"net/smtp"
	"untitled/internal/domain"
)

type SMTPConnector struct {
	Host string
	Port string
	From string
	User string
	Pass string
}

// Address URI to smtp server
func (s *SMTPConnector) Address() string {
	return s.Host + ":" + s.Port
}

func (s *SMTPConnector) Send(message domain.MailMessage) error {

	auth := smtp.PlainAuth("", s.User, s.Pass, s.Host)

	var body bytes.Buffer

	body.Write([]byte("To: " + message.To[0].String() + "\n"))
	body.Write([]byte("Subject: " + message.Subject + "\n"))
	body.Write([]byte("MIME-version: 1.0;\n"))
	body.Write([]byte("Content-Type: text/html; charset=utf-8; \n"))

	body.Write(message.Content)

	var to []string

	for _, tmp := range message.To {
		to = append(to, tmp.Address)
	}

	err := smtp.SendMail(s.Address(), auth, s.From, to, body.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	return err
}
