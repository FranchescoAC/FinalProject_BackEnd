package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
)

// GetUsers - Controller to fetch all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(w, "❌ Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
