package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv" // ✅ Importamos godotenv
	_ "github.com/lib/pq"
)

// ConnectDB establece una conexión a la base de datos usando variables de entorno.
func ConnectDB() (*sql.DB, error) {
	// ✅ Cargar las variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Advertencia: No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}

	// Obtener las variables de entorno
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Validar que las variables no estén vacías
	if host == "" || portStr == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("❌ Error: Faltan variables de entorno. Verifica tu archivo .env")
	}

	// Convertir el puerto a entero
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("❌ Error al convertir DB_PORT a entero:", err)
		return nil, err
	}

	// Construir la cadena de conexión
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Error conectando a la base de datos:", err)
		return nil, err
	}

	log.Println("✅ Conexión exitosa a la base de datos")
	return db, nil
}
