package routes

import (
	"userManagement/controller"
	"userManagement/middleware"
	"net/http"
	"github.com/gorilla/mux"
)

// SetRoutes - Configurar la ruta para actualizar usuarios
func SetRoutes(r *mux.Router) {
	r.Handle("/updateUser", middleware.JWTMiddleware(http.HandlerFunc(controller.UpdateUser))).Methods("PUT")
}

