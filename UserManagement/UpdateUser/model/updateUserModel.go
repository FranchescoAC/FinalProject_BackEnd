package model

import (
	"fmt"
	"userManagement/config"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUser - Actualizar los datos del usuario en la base de datos
func UpdateUser(updatedUser User) error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
	_, err = db.Exec(query, updatedUser.Username, updatedUser.Email, updatedUser.Password, updatedUser.ID)
	if err != nil {
		return fmt.Errorf("❌ Error actualizando usuario: %v", err)
	}

	return nil
}
