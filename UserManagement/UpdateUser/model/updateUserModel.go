package model

import (
	"fmt"
	"userManagement/config"
)

// User - Estructura de usuario (con los campos que necesitas)
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUser - Función para actualizar un usuario en la base de datos
func UpdateUser(updatedUser User) error {
	// Conectar a la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Crear la consulta SQL para actualizar los datos del usuario
	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`

	// Ejecutar la consulta
	_, err = db.Exec(query, updatedUser.Username, updatedUser.Email, updatedUser.Password, updatedUser.ID)
	if err != nil {
		return fmt.Errorf("Error updating user: %v", err)
	}

	return nil
}
