package main

import (
	"fmt"
	"net/http"
	"registerUser/config"
	"registerUser/routes"
<<<<<<< HEAD
	"github.com/rs/cors"
)

func main() {
	// Initialize the database connection and setup
	config.ConnectDatabase()
	config.InitDatabase()

	// Set up the routes for the RegisterUser microservice
	routes.RegisterRoutes()

	// Enable CORS for all routes and start the server
	r := http.NewServeMux()
	handler := cors.Default().Handler(r)

	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", handler)
}
=======
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
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
