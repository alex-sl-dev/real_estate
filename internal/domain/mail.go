package domain

type Mail struct {
	from string
	to []string
	subject string
	message []byte
}

type smtpServer struct {
	host string
	port string
}
// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

type MailService interface {
	Send (mail Mail) error
}