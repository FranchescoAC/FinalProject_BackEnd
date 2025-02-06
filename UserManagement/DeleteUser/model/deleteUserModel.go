package model

import (
	"fmt"
	"userManagement/config"
)

// DeleteUser - Deletes a user from the database by ID
func DeleteUser(userID int) error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM users WHERE id = $1`
	result, err := db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("❌ Error deleting user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("❌ Error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("❌ No user found with the provided ID")
	}

	return nil
}
