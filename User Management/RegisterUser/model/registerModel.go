package model

import (
	"registerUser/config"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

