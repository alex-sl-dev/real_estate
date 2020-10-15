package usecase

import "untitled/internal/domain"

type AccountService struct {
	AccountRepository domain.AccountRepository
	JWTService
	MailService /*maybe*/
}

func (service *AccountService) Register(aggregate domain.AccountAggregate) error {
	aggregate.Identity.EncryptPassword()
	err := service.AccountRepository.InsertAccount(aggregate)
	if err != nil {
		return err
	}
	// TODO temporary, broken single responsibility?
	// service.MailService.SendRegistrationMail(aggregate)
	return nil
}

func (service *AccountService) Authenticate(aggregate domain.AccountAggregate) (string, error) {
	aggregate.Identity.EncryptPassword()
	aggregate, err := service.AccountRepository.SelectAccountByCredentials(aggregate)
	if err != nil {
		return "", err
	}
	/*
	 we try to select by hashed password
		h := crypto.Hash{}

		err = h.Compare(u.Password, payload.Password)
		if err != nil {
			return nil, err
		}*/
	//
	token, err := service.JWTService.ClaimJWToken(aggregate)
	if err != nil {
		return "", err
	}

	return token, nil
}

