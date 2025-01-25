package routes

import (
	"ViewUser/controller"
	"github.com/gin-gonic/gin"
)

func ViewRoutes(router *gin.Engine) {
	viewController := controller.ViewController{}

	// Ruta para obtener todos los usuarios
	router.GET("/users", viewController.GetAllUsers)

	// Ruta existente para obtener un usuario por ID
	router.GET("/users/:id", viewController.GetUser)
}
