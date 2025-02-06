package main

import (
	"log"
	"net/http"
	"userManagement/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Crear el enrutador
	r := mux.NewRouter()

	// Registrar las rutas del microservicio
	routes.RegisterRoutes(r)

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir cualquier dominio (puedes restringirlo si lo deseas)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Envolver el router con el middleware de CORS
	handler := corsHandler.Handler(r)

	// Iniciar el servidor
	port := ":2001"
	log.Println("✅ Servidor de LoginUser iniciado en http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
