package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// ConnectDB establece una conexión a la base de datos y la retorna
func ConnectDB() (*sql.DB, error) {
	connStr := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("❌ Error conectando a la base de datos:", err)
		return nil, err
	}
	return db, nil
}
