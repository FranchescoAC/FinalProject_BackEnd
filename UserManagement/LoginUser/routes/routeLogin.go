package routes

import (
	"userManagement/controller"
	"github.com/gorilla/mux"
)

// Registrar la ruta del login
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")
}
