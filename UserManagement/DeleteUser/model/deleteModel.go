package model

import (
	"DeleteUser/config"
)

func DeleteUserByID(id string) error {
	db := config.GetDB()
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
