package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := GetRedis(r.Context()).Get(r.Header.Get("token"))
		if res.Err() != nil {
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
