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
	r.Use(
		middlewares.Database(api.GetDatabase()),
		middlewares.Redis(api.GetRedis()),
		middlewares.Authenticate,
	)
	r.Mount("/auth", auth.Routes())
	r.Mount("/customer", customer.Routes())
	log.Fatal(http.ListenAndServe(":8080", r))
}
