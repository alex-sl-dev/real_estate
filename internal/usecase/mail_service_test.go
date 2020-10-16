package usecase

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
	"untitled/internal/domain"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func TestMailService_SendRegistrationMail(t *testing.T) {

	mailServiceInstance := NewMailService(false)

	fullName := domain.FullNameValueObject{
		FirstName: "John",
		LastName:  "Doe",
	}

	profile := domain.ProfileValueObject{
		FullName: fullName,
	}

	email := domain.EmailValueObject{
		Value: "john.doe@localhost",
		//Value: "alex.slobodianiuk.84@gmail.com",
	}

	identity := domain.IdentityEntity{
		Email: email,
	}

	accountAggregate := domain.AccountAggregate{
		Profile:  profile,
		Identity: identity,
	}

	err := mailServiceInstance.SendRegistrationMail(accountAggregate)
	if err != nil {
		t.Error(err)
	}
}
