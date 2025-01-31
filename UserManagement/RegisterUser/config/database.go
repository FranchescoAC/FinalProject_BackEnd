package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	connStr := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	if err = DB.Ping(); err != nil {
		panic(fmt.Sprintf("Database is not reachable: %v", err))
	}

	fmt.Println("Database connected successfully!")
}

func InitDatabase() {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(query)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize database: %v", err))
	}
}