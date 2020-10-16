package usecase

import (
	"fmt"
	"testing"
)

var accountServiceInstance *AccountService

func init() {

	accountServiceInstance = &AccountService{
		//AccountRepository: accountRepositoryInstance,
		// MailService: *mailServiceInstance, /** todo, select proper place for service call service */
	}
}
func TestAccountService_GetVerificationCode(t *testing.T) {
	code := accountServiceInstance.GetVerificationCode("john.doe@localhost")
	if len(code) < 1 {
		t.Error("wrong code")
	}
	fmt.Println(code)
}
