package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
	"userManagement/service"
)

// Controlador para registrar un nuevo usuario
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User

	// Decodificamos el JSON del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Crear el usuario en la base de datos
	createdUser, err := model.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Enviar el webhook de forma asíncrona
	go service.SendWebhookToTicketManagement(createdUser)

	// Respondemos con un código 201 y los datos del nuevo usuario
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
