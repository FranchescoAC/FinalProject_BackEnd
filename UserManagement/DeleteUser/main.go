package main

import (
	"github.com/gin-gonic/gin"
	"DeleteUser/config"
	"DeleteUser/routes"
)

func main() {
	// Initialize the database connection and defer its closure
	config.InitDB()
	defer config.GetDB().Close()

	// Set up the routes
	r := gin.Default()
	routes.DeleteRoutes(r)

	// Start the server on port 8085
	r.Run(":8085")
}
