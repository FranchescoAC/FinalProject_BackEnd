package model

import (
	"database/sql"
	"userManagement/config"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UpdateUserByID(id string, user User) (*User, error) {
	db := config.GetDB()
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email"
	row := db.QueryRow(query, user.Name, user.Email, id)
	var updatedUser User
	if err := row.Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &updatedUser, nil
}
