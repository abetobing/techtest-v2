package auth

import "github.com/go-chi/chi"

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/login", loginHandler)
	return r
}
