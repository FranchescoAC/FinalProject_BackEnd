package main

import (
	"github.com/gin-gonic/gin"
	"LoginUser/config"
	"LoginUser/routes"
	"github.com/gin-contrib/cors"
)

func main() {
	// Initialize the database
	config.InitDB()
	defer config.GetDB().Close()

	// Create the Gin router and apply CORS middleware
	r := gin.Default()
	r.Use(cors.Default()) // Apply default CORS settings

	// Register the login routes
	routes.LoginRoutes(r)

	// Run the server
	r.Run(":8082")
}
