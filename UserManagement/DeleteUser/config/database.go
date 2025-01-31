package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	dsn := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}
}

func GetDB() *sql.DB {
	return db
}