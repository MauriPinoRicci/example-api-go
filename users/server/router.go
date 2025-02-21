package server

import (
	"github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/users", create)
	router.Get("/users/{id}", GetByID)
	router.Put("/users/{id}", Update)

	return router
}
