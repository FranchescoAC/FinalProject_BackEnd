package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importamos el driver de PostgreSQL
)

const (
	DbHost     = "userdb.ckkft1kjllti.us-east-1.rds.amazonaws.com"
	DbPort     = 5432
	DbUser     = "postgres"
	DbPassword = "Baseuser0806"
	DbName     = "UserManagement"
)

// Conectar a la base de datos PostgreSQL
func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", DbHost, DbPort, DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos: ", err)
		return nil, err
	}
	// Probar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Error al hacer ping a la base de datos: ", err)
		return nil, err
	}
	fmt.Println("✅ Conexión exitosa a la base de datos")
	return db, nil
}
