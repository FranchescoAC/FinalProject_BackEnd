package main

import (
	"github.com/gin-gonic/gin"
	"UpdateUser/config"
	"UpdateUser/routes"
)

func main() {
	// Initialize the database connection and defer its closure
	config.InitDB()
	defer config.GetDB().Close()

	// Set up the routes and CORS
	r := gin.Default()
	routes.UpdateRoutes(r)

	// Start the server on port 8084
	r.Run(":8084")
}
