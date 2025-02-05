package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
)

// UpdateUser - Controlador para actualizar los datos de un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Crear una estructura para los datos del usuario a actualizar
	var updatedUser model.User

	// Decodificar el cuerpo de la solicitud JSON
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llamar a la función de modelo para actualizar el usuario en la base de datos
	err = model.UpdateUser(updatedUser)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	// Devolver una respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated successfully",
	})
}
