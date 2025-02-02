package routes

import (
	"userManagement/controller"
	"github.com/gorilla/mux"
)

// Función para registrar las rutas de LoginUser
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")
}
