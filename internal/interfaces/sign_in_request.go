package interfaces

import "untitled/internal/domain"

type SignInRequest struct {
	Email     string    `validate:"required,email" json:"email"`
	Password  string    `validate:"required,ascii" json:"password"`
}

func (r SignInRequest) ToAccountAggregate() domain.AccountAggregate {

	accountAggregate := domain.AccountAggregate{
		Identity: domain.IdentityEntity{},
	}

	accountAggregate.Identity.Email.Value = r.Email
	accountAggregate.Identity.Password = r.Password

	return accountAggregate
}
