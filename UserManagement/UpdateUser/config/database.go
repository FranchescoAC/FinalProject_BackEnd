package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Importamos el driver de PostgreSQL
)

// ConnectDB - Función para establecer la conexión a la base de datos
func ConnectDB() (*sql.DB, error) {
	// Cadena de conexión con los datos proporcionados
	connStr := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
	
	// Abrimos la conexión con PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error al conectar a la base de datos: %v", err)
	}

	// Verificamos si la conexión es exitosa
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("No se pudo hacer ping a la base de datos: %v", err)
	}

	// Si la conexión es exitosa, devolvemos la instancia de la conexión
	return db, nil
}
