package model

import (
	"database/sql"
	"ViewUser/config"
	"fmt"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Obtener un usuario por su ID
func GetUserByID(id string) (*User, error) {
	// Convertir el ID a entero
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}

	db := config.GetDB()
	query := "SELECT id, name, email FROM users WHERE id = $1"
	row := db.QueryRow(query, userID)

	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, err
	}
	return &user, nil
}

// Obtener todos los usuarios
func GetAllUsers() ([]User, error) {
	db := config.GetDB()
	query := "SELECT id, name, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Comprobar errores de iteración
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
