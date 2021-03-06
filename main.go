package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"customer/api"
	"customer/api/auth"
	"customer/api/customer"
	"customer/api/middlewares"
)

func main() {
	err := api.LoadConfig(nil)
	if err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()

	// Guest: No authentication
	r.Group(func(r chi.Router) {
		r.Use(
			middlewares.Database(api.GetDatabase()),
			middlewares.Redis(api.GetRedis()),
			middlewares.Header,
			middlewares.Cors(),
		)
		r.Mount("/auth", auth.Routes())
	})

	// Resource: with auth
	r.Group(func(r chi.Router) {
		r.Use(
			middlewares.Database(api.GetDatabase()),
			middlewares.Redis(api.GetRedis()),
			middlewares.Header,
			// middlewares.Authenticate, // disable old token authentication, now using jwt
			middlewares.Jwt,
			middlewares.Cors(),
		)
		r.Mount("/customer", customer.Routes())
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
