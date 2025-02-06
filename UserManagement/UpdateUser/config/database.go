package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Importamos el driver de PostgreSQL
)

// ConnectDB - Conectar a la base de datos
func ConnectDB() (*sql.DB, error) {
	connStr := "host=userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com port=5432 user=postgres password=Baseuser0806 dbname=UserManagement sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("❌ Error al conectar a la base de datos: %v", err)
	}

	// Verificar conexión
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("❌ No se pudo hacer ping a la base de datos: %v", err)
	}

	fmt.Println("✅ Conexión exitosa a la base de datos")
	return db, nil
}
