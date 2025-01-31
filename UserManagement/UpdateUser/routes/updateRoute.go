package routes

import (
	"net/http"  // Agregar esta importación
	"github.com/gin-gonic/gin"
	"UpdateUser/controller"
	"github.com/rs/cors"
)

func UpdateRoutes(router *gin.Engine) {
	updateController := controller.UpdateController{}

	// Definir la ruta para actualizar un usuario
	router.PUT("/update/:id", updateController.UpdateUser)

	// Habilitar CORS para todas las rutas
	handler := cors.Default().Handler(router)

	// Usar el handler para envolver las rutas con el middleware CORS
	// Iniciar el servidor usando el handler con CORS habilitado
	http.ListenAndServe(":8084", handler)  // Aquí usa http.ListenAndServe con el handler
}
