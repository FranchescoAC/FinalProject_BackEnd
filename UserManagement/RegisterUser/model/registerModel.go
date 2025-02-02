package model;

import (
	"database/sql"         // Asegúrate de que esta línea está incluida
	"userManagement/config" // Conexión a la base de datos
	_ "github.com/lib/pq"   // Controlador PostgreSQL
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Función para crear un nuevo usuario en la base de datos
func CreateUser(user User) (User, error) {
	// Conectamos con la base de datos usando el archivo config
	db, err := config.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close() // Cerramos la conexión al final

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

// Agrega una función que use explícitamente el paquete `database/sql`
func ExampleSQLUsage() *sql.DB {
	db, _ := sql.Open("postgres", "connection_string")
	return db
}