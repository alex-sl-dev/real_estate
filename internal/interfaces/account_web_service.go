package interfaces

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"untitled/internal/domain"
)

type AccountWebService struct {
	AccountService domain.AccountService
}

func (aws *AccountWebService) SignUpAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var (
		registrationRequest RegistrationRequest
		err                 error
	)

	err = json.NewDecoder(r.Body).Decode(&registrationRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Validate.Struct(&registrationRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = aws.AccountService.Register(registrationRequest.ToAccountAggregate())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("201 - User Account Registered!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (aws *AccountWebService) SignInAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		signInRequest SignInRequest
		err           error
	)

	err = json.NewDecoder(r.Body).Decode(&signInRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Validate.Struct(&signInRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	token, err := aws.AccountService.Authenticate(signInRequest.ToAccountAggregate())
	if err != nil && err.Error() == "no rows in result set" {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := SignInResponse{
		Message: "Authorized",
		Token:   token,
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/jsonResponse")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (aws *AccountWebService) SignOutAction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// We don't using user session, here we do nothing.
	// Just pass for sign out action, then on client side, app should forget JWT
	var err error
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("200 - Logout"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (aws *AccountWebService) LoadProfileAction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {


}

func (aws *AccountWebService) UpdateProfileAction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {


}

