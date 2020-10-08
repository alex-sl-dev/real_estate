package interfaces

import (
	"time"
	"untitled/internal/domain"
)

type SignInResponse struct {
	Message string          `json:"message"`
	Account AccountResponse `json:"account"`
	Token   string          `json:"token"`
}

type AccountResponse struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	PostIndex string    `json:"post_index"`
	AboutMe   string    `json:"about_me"`
	Avatar    string    `json:"avatar"`
	CreatedAd time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func AccountResponseFromAccountAggregate(aggregate domain.AccountAggregate) (AccountResponse, error) {

	return AccountResponse{}, nil
}
