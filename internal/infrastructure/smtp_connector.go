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

	body.Write([]byte("Subject:" + message.Subject + "\n"))
	body.Write([]byte("MIME-version: 1.0;\n"))
	body.Write([]byte("Content-Type: text/html;\n"))

	body.Write(message.Content)

	err := smtp.SendMail(s.Address(), auth, s.From, message.To, body.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	return err

	/*
	// Connect to the remote SMTP server.
	c, err := smtp.Dial()

	// Set the sender and recipient first
	if err := c.Mail(s.From); err != nil {
		log.Fatal(err)
	}

	for _, to := range message.To {
		if err := c.Rcpt(to); err != nil {
			log.Fatal(err)
		}
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	var headers []string

	headers = append(headers, "Subject:" + message.Subject + "\n")
	headers = append(headers, "MIME-version: 1.0;\n")
	headers = append(headers, "Content-Type: text/html;\n")

	_, err = fmt.Fprintf(wc, strings.Join(headers, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(wc, message.Content)
	if err != nil {
		log.Fatal(err)
	}

	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
	 */

}