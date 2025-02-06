package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

var db *sql.DB

// ConnectDB - Establish a connection to the PostgreSQL database
func ConnectDB() (*sql.DB, error) {
	connStr := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("❌ Error connecting to the database: %v", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("❌ Could not establish a connection: %v", err)
	}

	fmt.Println("✅ Successfully connected to the database")
	return db, nil
}
