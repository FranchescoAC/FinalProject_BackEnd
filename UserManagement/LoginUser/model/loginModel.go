package model

import (
	"userManagement/config" // Mantén solo la importación de config
	_ "github.com/lib/pq"    // Importamos el driver para PostgreSQL
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Función para obtener un usuario por correo electrónico
func GetUserByEmail(email string) (User, error) {
	// Usamos la función ConnectDB del paquete config para obtener la conexión
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
