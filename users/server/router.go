package server

import (
	"github.com/go-chi/chi/v5"
)

func InitRouter(dep *Dependencies) *chi.Mux {
	router := chi.NewRouter()

	handlers := NewHandlers(dep)

	router.Post("/users", handlers.create)
	router.Get("/users/{id}", handlers.GetByID)
	router.Put("/users/{id}", handlers.Update)
	router.Delete("/users/{id}", handlers.Delete)

	return router
}
