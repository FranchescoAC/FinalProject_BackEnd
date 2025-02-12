package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// ConnectDB establece una conexión a la base de datos y la retorna usando variables de entorno.
func ConnectDB() (*sql.DB, error) {
	// Se obtienen las variables de entorno
	host := os.Getenv("DB_HOST")         
	portStr := os.Getenv("DB_PORT")        
	user := os.Getenv("DB_USER")           
	password := os.Getenv("DB_PASSWORD")   
	dbname := os.Getenv("DB_NAME")         

	// Convertir el puerto a entero
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Println("❌ Error al convertir DB_PORT a entero:", err)
		return nil, err
	}

	// Construir la cadena de conexión
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("❌ Error conectando a la base de datos:", err)
		return nil, err
	}

	return db, nil
}
