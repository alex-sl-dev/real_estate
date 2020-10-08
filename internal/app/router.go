package app

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strings"
)

func gwtAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("GWT_KEY")), nil
			})

			fmt.Println(token.Claims.(jwt.MapClaims))
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next(w, r.WithContext(ctx), ps)
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	}
}

// NewRouter provide router definition
func NewRouter() *httprouter.Router {
	router := httprouter.New()


	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	router.ServeFiles("/static/*filepath", http.Dir(pwd+"/public/"))


	accountWebService := sm.AccountWebServiceFactory()
	router.POST("/account/sign-up", accountWebService.SignUpAction)
	router.POST("/account/sign-in", accountWebService.SignInAction)
	router.POST("/account/sign-out", accountWebService.SignOutAction)

	router.GET("/account/profile", gwtAuth(accountWebService.LoadProfileAction))
	router.POST("/account/profile", gwtAuth(accountWebService.UpdateProfileAction))

	return router
}
