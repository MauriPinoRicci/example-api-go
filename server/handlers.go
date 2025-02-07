package server

import (
	"net/http"
)

func getMensaje(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("¡Hola desde Chi con un Controller! 🚀"))
}
