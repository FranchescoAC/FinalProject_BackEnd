package routes

import (
	"userManagement/controller"
	"github.com/gorilla/mux"
)

// SetRoutes - Función para configurar las rutas
func SetRoutes(r *mux.Router) {
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
}
