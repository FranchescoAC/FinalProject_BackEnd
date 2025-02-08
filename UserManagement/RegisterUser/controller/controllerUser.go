package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
	"userManagement/service"
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User

	// Decode the JSON request body
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create the user in the database
	createdUser, err := model.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Send the webhook asynchronously to NotificationManagement
	go service.SendWebhookToNotificationManagement(createdUser)

	// Respond with the newly created user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
