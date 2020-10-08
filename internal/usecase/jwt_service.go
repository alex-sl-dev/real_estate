package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
	"untitled/internal/domain"
)

type JWTService struct {

}

func (service *JWTService) ClaimJWToken(aggregate domain.AccountAggregate) (string, error) {
	// Create a struct that will be encoded to a JWT.
	// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
	type Claims struct {
		UID string `json:"uid"`
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := Claims{
		UID: strconv.Itoa(int(aggregate.Identity.ID)),
		Foo: "bar",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("GWT_KEY")))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}

	return tokenString, nil
}
