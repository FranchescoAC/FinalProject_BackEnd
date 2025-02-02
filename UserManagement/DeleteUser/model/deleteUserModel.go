package model

import (
	"fmt"
	"userManagement/config"
)

// DeleteUser - Función para eliminar un usuario de la base de datos
func DeleteUser(userID int) error {
	// Conectar a la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Crear la consulta SQL para eliminar el usuario
	query := `DELETE FROM users WHERE id = $1`

	// Ejecutar la consulta
	result, err := db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("Error al eliminar el usuario: %v", err)
	}

	// Verificar si se eliminó un usuario
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error al verificar la eliminación: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No se encontró un usuario con el ID proporcionado")
	}

	return nil
}
