package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
)

// GetUsers - Controlador para obtener todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Obtener todos los usuarios desde la base de datos
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
		return
	}

	// Devolver los usuarios en formato JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
