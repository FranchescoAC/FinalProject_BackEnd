package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
	"userManagement/service"
)

// Controlador para manejar el inicio de sesión del usuario
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Decodificar la solicitud JSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Verificar que el usuario exista en la base de datos
	storedUser, err := model.GetUserByEmail(user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Verificar que la contraseña sea válida (en producción, se debería hashear)
	if storedUser.Password != user.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generar el token JWT
	tokenString, err := service.GenerateJWT(storedUser.ID, storedUser.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Responder con el token JWT
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
