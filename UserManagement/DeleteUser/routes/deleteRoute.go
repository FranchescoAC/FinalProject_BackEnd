package routes

import (
	"userManagement/controller"
	"github.com/gorilla/mux"
)

// SetRoutes - Configuración de rutas para DeleteUser
func SetRoutes(r *mux.Router) {
	r.HandleFunc("/deleteUser", controller.DeleteUser).Methods("DELETE")
}
