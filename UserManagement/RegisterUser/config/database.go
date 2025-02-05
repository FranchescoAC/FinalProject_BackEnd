package config

import (
	"database/sql"
	_ "github.com/lib/pq" // El paquete pq es necesario para interactuar con PostgreSQL
)

// ConnectDB establishes a connection to the database and returns a database object.
func ConnectDB() (*sql.DB, error) {
	connStr := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
