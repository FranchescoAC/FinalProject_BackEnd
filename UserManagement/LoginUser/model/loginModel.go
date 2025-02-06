package model

import (
	"userManagement/config"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Obtener un usuario por su correo electrónico
func GetUserByEmail(email string) (User, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	var user User
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	err = db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
