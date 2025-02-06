package main

import (
	"log"
	"net/http"
	"userManagement/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Crear el router
	r := mux.NewRouter()

	// Configurar las rutas
	routes.SetRoutes(r)

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permite cualquier origen (puedes restringirlo)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Envolver el router con CORS
	handler := corsHandler.Handler(r)

	// Iniciar el servidor
	port := ":2003"
	log.Println("✅ Servidor de UpdateUser corriendo en http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
