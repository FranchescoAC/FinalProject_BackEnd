package main

import (
	"github.com/gin-gonic/gin"
	"LoginUser/config"
	"LoginUser/routes"
<<<<<<< HEAD
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
=======
)

func main() {
	config.InitDB()
	defer config.GetDB().Close()

	r := gin.Default()
	routes.LoginRoutes(r)

	r.Run(":8082")
}
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
