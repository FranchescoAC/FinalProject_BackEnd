package model

import (
	"userManagement/config"
)

// Función para crear un nuevo usuario en la base de datos
func CreateUser(user User) (User, error) {
	// Conectamos con la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	// Query SQL para insertar un nuevo usuario
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	var userID int
	err = db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return User{}, err
	}

	// Asignamos el ID generado a la estructura User
	user.ID = userID
	return user, nil
}
