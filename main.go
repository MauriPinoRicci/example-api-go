package main

import (
	"fmt"
	"net/http"

	"github.com/MauriPinoRicci/example-api-go/server"
)

func main() {

	router := server.InitRouter()

	// Levantar el servidor
	port := ":8080"
	fmt.Println("Servidor corriendo en http://localhost" + port)
	http.ListenAndServe(port, router)
}
