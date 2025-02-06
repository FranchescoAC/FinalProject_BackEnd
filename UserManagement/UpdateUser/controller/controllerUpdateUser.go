package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
)

// UpdateUser - Controlador para actualizar los datos del usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser model.User

	// Decodificar JSON de la solicitud
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "❌ Entrada inválida", http.StatusBadRequest)
		return
	}

	// Llamar al modelo para actualizar el usuario
	if err := model.UpdateUser(updatedUser); err != nil {
		http.Error(w, "❌ Error actualizando usuario", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Usuario actualizado exitosamente",
	})
}
