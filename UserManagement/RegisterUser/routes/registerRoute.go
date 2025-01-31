package routes

import (
	"net/http"
	"registerUser/controller"
	"registerUser/utils"
)

func RegisterRoutes() {
	// Ruta pública: Registro de usuario
	http.HandleFunc("/register", controller.RegisterUser)
	// Ruta pública: Inicio de sesión
	http.HandleFunc("/login", controller.LoginUser)

	// Rutas protegidas por JWT
	http.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Validar el token JWT antes de proceder
		_, err := utils.ValidateJWT(w, r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Si el token es válido, acceder a los datos del usuario
		controller.ViewUser(w, r)
	})

	http.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		// Validar el token JWT antes de proceder
		_, err := utils.ValidateJWT(w, r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Si el token es válido, proceder a actualizar los datos
		controller.UpdateUser(w, r)
	})

	http.HandleFunc("/user/delete", func(w http.ResponseWriter, r *http.Request) {
		// Validar el token JWT antes de proceder
		_, err := utils.ValidateJWT(w, r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Si el token es válido, proceder a eliminar el usuario
		controller.DeleteUser(w, r)
	})
}
