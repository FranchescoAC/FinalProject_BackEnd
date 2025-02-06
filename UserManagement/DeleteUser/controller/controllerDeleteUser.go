package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"userManagement/model"
)

// DeleteUser - Controller to delete a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL query parameters
	userIDStr := r.URL.Query().Get("id")
	if userIDStr == "" {
		http.Error(w, "❌ User ID is required", http.StatusBadRequest)
		return
	}

	// Convert the user ID from string to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "❌ Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the model function to delete the user
	if err := model.DeleteUser(userID); err != nil {
		http.Error(w, "❌ Error deleting user", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ User successfully deleted",
	})
}
