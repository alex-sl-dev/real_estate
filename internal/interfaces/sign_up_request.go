package interfaces

import (
	"time"
	"untitled/internal/domain"
)

// RegistrationRequest entity
type RegistrationRequest struct {
	ID        int64     `validate:"numeric" json:"id"`
	FirstName string    `validate:"required,alpha" json:"first_name"`
	LastName  string    `validate:"required,alpha" json:"last_name"`
	Email     string    `validate:"required,email" json:"email"`
	Password  string    `validate:"required,ascii" json:"password"`
	Phone     string    `validate:"required,numeric" json:"phone"`
	Role      string    `validate:"required,numeric" json:"role"`
	Company   string    `validate:"omitempty,alpha" json:"company"`
	Address   string    `validate:"omitempty,alpha" json:"address"`
	City      string    `validate:"omitempty,alpha" json:"city"`
	Country   string    `validate:"omitempty,alpha" json:"country"`
	PostIndex string    `validate:"omitempty,numeric" json:"post_index"`
	AboutMe   string    `validate:"omitempty,alpha" json:"about_me"`
	Avatar    string    `validate:"omitempty,file" json:"avatar"`
	CreatedAd time.Time `validate:"omitempty,numeric" json:"created_at"`
	UpdatedAt time.Time `validate:"omitempty,numeric" json:"updated_at"`
	DeletedAt time.Time `validate:"omitempty,numeric" json:"deleted_at"`
}

func (r *RegistrationRequest) ToAccountAggregate() domain.AccountAggregate {

	accountAggregate := domain.AccountAggregate{
		Identity: domain.IdentityEntity{},
		Profile: domain.ProfileValueObject{
			FullName: domain.FullNameValueObject{},
			Phone:    domain.PhoneValueObject{},
		},
		Address:  domain.AddressValueObject{},
	}

	accountAggregate.Identity.Email.Value = r.Email
	accountAggregate.Identity.Password = r.Password
	accountAggregate.Identity.Role = r.Role

	accountAggregate.Profile.FullName.FirstName = r.FirstName
	accountAggregate.Profile.FullName.LastName = r.LastName

	//accountAggregate.Phone.Number

	return accountAggregate
}

