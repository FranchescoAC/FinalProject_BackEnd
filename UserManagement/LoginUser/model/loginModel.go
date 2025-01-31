package model

import (
	"LoginUser/config"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Authenticate() bool {
	db := config.GetDB()
	query := "SELECT COUNT(*) FROM users WHERE email=$1 AND password=$2"
	row := db.QueryRow(query, u.Email, u.Password)
	var count int
	if err := row.Scan(&count); err != nil {
		return false
	}
	return count > 0
}