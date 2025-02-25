package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MauriPinoRicci/example-api-go/users/server"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar archivo .env: %v", err)
	}

	// Inicializamos las dependencias
	dependencies := server.InitDependencies()

	// Inicializamos el servidor
	router := server.InitRouter(dependencies)

	// Levantar el servidor
	port := ":8080"
	fmt.Println("Servidor corriendo en http://localhost" + port)
	http.ListenAndServe(port, router)
}
