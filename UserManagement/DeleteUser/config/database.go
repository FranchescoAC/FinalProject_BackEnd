package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL
)

var db *sql.DB

// ConnectDB - Establece la conexión a la base de datos PostgreSQL usando variables de entorno.
func ConnectDB() (*sql.DB, error) {
	// Se leen las variables de entorno
	host := os.Getenv("DB_HOST")         
	portStr := os.Getenv("DB_PORT")       
	user := os.Getenv("DB_USER")           
	password := os.Getenv("DB_PASSWORD")   
	dbname := os.Getenv("DB_NAME")         

	// Convertir el puerto a entero
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("❌ Error al convertir DB_PORT a entero: %v", err)
	}

	// Construir la cadena de conexión
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	// Abrir la conexión
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("❌ Error connecting to the database: %v", err)
	}

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("❌ Could not establish a connection: %v", err)
	}

	fmt.Println("✅ Successfully connected to the database")
	return db, nil
}
