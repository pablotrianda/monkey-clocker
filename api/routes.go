package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes() *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Get("/", sayHello)
	r.Post("/schema", newSchema)
	r.Get("/{id}", wichID)

	return r
}
