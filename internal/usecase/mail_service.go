package usecase

import "untitled/internal/domain"

type MailService struct {

}


func (ms *MailService) Send(mail domain.Mail) error  {
	return nil
}
