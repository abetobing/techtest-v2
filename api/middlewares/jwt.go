package middlewares

import (
	"encoding/json"
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
			// fmt.Fprintln(w, "Unauthorize. Please login!")
			// w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusUnauthorized)

			generalError := GeneralError{Error: http.StatusText(http.StatusUnauthorized)}
			e, err := json.Marshal(generalError)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(e)
			return
		}

		next.ServeHTTP(w, r)
	})
}
