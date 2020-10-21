package app

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	NewServiceContainer()
}

func isProductionEnv() bool {
	env := os.Getenv("ENV")
	isProduction := false
	if env == "prod" || env == "production" {
		isProduction = true
	}
	return isProduction
}

func extractClaims(r *http.Request) (jwt.MapClaims, error) {
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

	if len(authHeader) != 2 {
		return nil, errors.New("invaid JWT token")
	} else {
		hmacSecretString := os.Getenv("GWT_KEY") // Value
		hmacSecret := []byte(hmacSecretString)
		token, err := jwt.Parse(authHeader[1], func(token *jwt.Token) (interface{}, error) {
			// check token signing method etc
			return hmacSecret, nil
		})

		if err != nil {
			return nil, err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims, nil
		} else {
			log.Printf("Invalid JWT Token")
			return nil, errors.New("invaid JWT token")
		}
	}
}

// GetFileContentType helper
func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
