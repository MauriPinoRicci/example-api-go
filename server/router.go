package server

import (
	"github.com/go-chi/chi/v5"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()

	// Endpoint GET que devuelve un mensaje usando el controlador
	router.Get("/", getMensaje)
	

	return router
}
