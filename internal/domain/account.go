package domain

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AccountAggregate struct {
	Identity IdentityEntity
	Profile  ProfileValueObject
	Address  AddressValueObject
}

type IdentityEntity struct {
	ID        int
	Password  string
	Email     EmailValueObject
	Role      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (ie IdentityEntity) EncryptPassword() {
	saltedBytes := []byte(ie.Password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	ie.Password = string(hashedBytes[:])
}

type EmailValueObject struct {
	Value           string
	VerifiedAt      time.Time
	ActivationToken string
}

type ProfileValueObject struct {
	FullName FullNameValueObject
	Phone    PhoneValueObject
	AboutMe  string
	Avatar   string
	Company  string
}

type PhoneValueObject struct {
	FullPhone string
	Country   string
	Code      string
	Number    string
}

func (pvo *PhoneValueObject) getFullPhone() {
	pvo.FullPhone = pvo.Country + pvo.Code + pvo.Number
}

func (pvo *PhoneValueObject) FromString(fullPhone string) {
	if len(fullPhone) > 10 {
		pvo.FullPhone = fullPhone
		tmp := []rune(fullPhone)
		pvo.Country = string(tmp[0:2])
		pvo.Code = string(tmp[2:4])
		pvo.Number = string(tmp[4:len(fullPhone)])
	}
}

type FullNameValueObject struct {
	FirstName string
	LastName  string
}

func (o *FullNameValueObject) FullName() string {
	return o.FirstName + " " + o.LastName
}

type AddressValueObject struct {
	Country string
	City    string
	Post    string
}

type AccountRepository interface {
	SelectAccountByID(id int) (AccountAggregate, error)
	SelectAccountByCredentials(aggregate AccountAggregate) (AccountAggregate, error)
	InsertAccount(aggregate AccountAggregate) error
	//UpdateAccount(aggregate AccountAggregate) error
}

// AccountService can be named as AccountInteractor in some scenarios
type AccountService interface {
	Register(aggregate AccountAggregate) error
	Authenticate(aggregate AccountAggregate) (string, error)
	GetVerificationCode(email string) string
	//Disable(aggregate AccountAggregate) error
}
