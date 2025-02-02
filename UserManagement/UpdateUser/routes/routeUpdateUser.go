package routes

import (
	"userManagement/controller"
	"github.com/gorilla/mux"
)

// SetRoutes - Configuración de rutas para UpdateUser
func SetRoutes(r *mux.Router) {
	r.HandleFunc("/updateUser", controller.UpdateUser).Methods("PUT")
}
