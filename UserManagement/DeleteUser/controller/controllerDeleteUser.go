package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
	"strconv"
)

// DeleteUser - Controlador para eliminar un usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario a eliminar desde los parámetros de la URL
	userIDStr := r.URL.Query().Get("id")
	if userIDStr == "" {
		http.Error(w, "ID del usuario es requerido", http.StatusBadRequest)
		return
	}

	// Convertir el ID de string a int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Llamar al modelo para eliminar el usuario
	err = model.DeleteUser(userID)
	if err != nil {
		http.Error(w, "Error eliminando usuario", http.StatusInternalServerError)
		return
	}

	// Devolver una respuesta exitosa
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Usuario eliminado exitosamente",
	})
}
