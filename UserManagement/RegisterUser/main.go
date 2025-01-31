package main

import (
	"fmt"
	"net/http"
	"registerUser/config"
	"registerUser/routes"
	"github.com/gorilla/handlers" // Importar el paquete de CORS
)

func main() {
	// Conexión y configuración de la base de datos
	config.ConnectDatabase()
	config.InitDatabase()

	// Registro de rutas
	routes.RegisterRoutes()

	// Configuración de CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})  // Permite todos los orígenes
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Inicia el servidor con CORS habilitado
	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", handlers.CORS(origins, headers, methods)(nil))
}