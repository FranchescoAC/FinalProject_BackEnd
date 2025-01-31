package routes

import (
<<<<<<< HEAD
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
=======
	"github.com/gin-gonic/gin"
	"UpdateUser/controller"
)

func UpdateRoutes(router *gin.Engine) {
	controller := controller.UpdateController{}
	router.PUT("/update/:id", controller.UpdateUser)
}
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
