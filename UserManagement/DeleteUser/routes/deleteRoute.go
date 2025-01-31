package routes

import (
	"github.com/gin-gonic/gin"
	"DeleteUser/controller"
	"github.com/rs/cors"
)

func DeleteRoutes(router *gin.Engine) {
	deleteController := controller.DeleteController{}

	// Route to delete a user by ID
	router.DELETE("/users/:id", deleteController.DeleteUser)

	// Enable CORS for all routes
	handler := cors.Default().Handler(router)

	// Start the server
	router.Run(":8085")
}
