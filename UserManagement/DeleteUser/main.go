package main

import (
	"log"
	"net/http"
	"userManagement/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Create the router
	r := mux.NewRouter()

	// Set up the routes
	routes.SetRoutes(r)

	// Configure CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins (you can restrict if needed)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS middleware
	handler := corsHandler.Handler(r)

	// Start the server on port 2004
	port := ":2004"
	log.Println("✅ DeleteUser service running at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
