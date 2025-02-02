package model

import (
	"userManagement/config"
)

// User - Estructura para los usuarios
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetAllUsers - Función para obtener todos los usuarios de la base de datos
func GetAllUsers() ([]User, error) {
	// Conectar a la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()  // Asegurarse de cerrar la conexión después

	// Hacer la consulta para obtener todos los usuarios
	rows, err := db.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Cerrar las filas después de usarlas

	// Crear un slice de usuarios
	var users []User
	for rows.Next() {
		var user User
		// Escanear los valores de la fila en la estructura User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user) // Agregar el usuario al slice
	}

	// Verificar si hubo algún error al iterar sobre las filas
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
