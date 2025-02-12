package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq" // Importamos el driver de PostgreSQL
)

// ConnectDB - Conectar a la base de datos usando variables de entorno
func ConnectDB() (*sql.DB, error) {
	// Obtener los valores desde las variables de entorno
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

	// Conectar a la base de datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("❌ Error al conectar a la base de datos: %v", err)
	}

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("❌ No se pudo hacer ping a la base de datos: %v", err)
	}

	fmt.Println("✅ Conexión exitosa a la base de datos")
	return db, nil
}
