package middlewares

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
)

type GeneralError struct {
	Error string `json:"error"`
}

func Jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
		jwtauth.Verifier(tokenAuth)

		tokenString := jwtauth.TokenFromHeader(r)
		res := GetRedis(r.Context()).Get(tokenString)
		if res.Err() != nil {
			log.Printf("JWT auth failed, fallback to old token authentication")
			Authenticate(next)
		}

		next.ServeHTTP(w, r)
	})
}
