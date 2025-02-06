package routes

import (
	"userManagement/controller"
	"github.com/gorilla/mux"
)

// Registrar las rutas
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", controller.RegisterUser).Methods("POST")
}




