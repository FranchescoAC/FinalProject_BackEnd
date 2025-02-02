package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

// ConnectDB - Función para establecer la conexión con la base de datos PostgreSQL
func ConnectDB() (*sql.DB, error) {
	if db == nil {
		var err error
		connStr := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			return nil, fmt.Errorf("Error al conectar a la base de datos: %v", err)
		}
	}
	return db, nil
}
