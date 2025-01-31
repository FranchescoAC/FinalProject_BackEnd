package main

import (
	"fmt"
	"net/http"
	"registerUser/config"
	"registerUser/routes"
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
