package routes

import (
	"github.com/gin-gonic/gin"
	"ViewUser/controller"
	"github.com/rs/cors"
)

func ViewRoutes(router *gin.Engine) {
	viewController := controller.ViewController{}

	// Route to get all users
	router.GET("/users", viewController.GetAllUsers)

	// Route to get a user by ID
	router.GET("/users/:id", viewController.GetUser)

	// Enable CORS for all routes
	handler := cors.Default().Handler(router)

	// Start the server
	router.Run(":8083")
}
