package routes

import (
	"github.com/gin-gonic/gin"
	"DeleteUser/controller"
<<<<<<< HEAD
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
=======
)

func DeleteRoutes(router *gin.Engine) {
	controller := controller.DeleteController{}
	router.DELETE("/delete/:id", controller.DeleteUser)
}
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
