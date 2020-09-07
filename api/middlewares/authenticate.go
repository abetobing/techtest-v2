package middlewares

import (
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := GetRedis(r.Context()).Get(r.Header.Get("token"))
		if res.Err() != nil {
			// fmt.Fprintln(w, "Unauthorize. Please login!")
			w.WriteHeader(http.StatusUnauthorized)
		}
		next.ServeHTTP(w, r)
	})
}
