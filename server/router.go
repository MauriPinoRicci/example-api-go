package server

import (
	"github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/users", createUser)
	

	return router
}
