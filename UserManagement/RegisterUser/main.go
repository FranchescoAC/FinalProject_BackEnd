package main

import (
	"log"
	"net/http"

	"userManagement/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// Registrar rutas
	routes.RegisterRoutes(r)

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(r)

	// Configurar puerto del servidor
	port := ":2000"
	log.Println("\u2705 Servidor iniciado en http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}