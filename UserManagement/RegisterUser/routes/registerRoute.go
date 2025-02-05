package routes

import (
	"userManagement/controller"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Definir las rutas de la API REST y asociarlas con las funciones del controlador
	r.HandleFunc("/register", controller.RegisterUser).Methods("POST")  // Ruta REST para registrar un nuevo usuario
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")        // Ruta REST para hacer login

	return r
}



