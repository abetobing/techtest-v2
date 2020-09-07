package customer

import "github.com/go-chi/chi"

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create", createHandler)
	r.Get("/get/{id}", getHandler)
	r.Get("/list", listHandler)
	r.Delete("/delete/{id}", deleteHandler)
	return r
}
