package main

import (
	"log"
	"net/http"
	"os"

	"userManagement/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load server port from environment variable (default: 2000)
	port := os.Getenv("PORT")
	if port == "" {
		port = "2000"
	}

	// Initialize router
	r := mux.NewRouter()

	// Register routes for user management
	routes.RegisterRoutes(r)

	// Configure CORS to allow external requests
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow requests from any origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Wrap router with CORS middleware
	handler := corsHandler.Handler(r)

	// Configure server to listen on 0.0.0.0:<port> for external access
	serverAddress := "0.0.0.0:" + port
	log.Println("✅ Server running at http://" + serverAddress)

	// Start HTTP server
	log.Fatal(http.ListenAndServe(serverAddress, handler))
}
