package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Importar el driver de PostgreSQL
)

var db *sql.DB

// ConnectDB - Función para establecer la conexión a la base de datos
func ConnectDB() (*sql.DB, error) {
	// Usa tu cadena de conexión aquí
	connStr := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
	
	// Establecer la conexión
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error al conectar a la base de datos: %v", err)
	}

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("No se pudo hacer ping a la base de datos: %v", err)
	}

	return db, nil
}

// GetDB - Función para obtener la conexión a la base de datos
func GetDB() (*sql.DB) {
	return db
}
