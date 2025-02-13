package config

import (
    "database/sql"
    "fmt"
    "os"
    "strconv"
    "time"

    _ "github.com/lib/pq"
)

// Conectar a la base de datos PostgreSQL usando variables de entorno
func ConnectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Convertir el puerto a entero
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("❌ Error al convertir DB_PORT a entero: %w", err) // Usar fmt.Errorf para envolver el error original
	}

	// Construir la cadena de conexión
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	var db *sql.DB
	retryAttempts := 3       // Número de reintentos
	retryDelay := time.Second * 5 // Tiempo de espera entre reintentos

	for i := 0; i < retryAttempts; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("✅ Conexión exitosa a la base de datos")
				return db, nil
			}
		}

		fmt.Printf("❌ Error al conectar a la base de datos (intento %d): %v\n", i+1, err)
		if i < retryAttempts-1 {
			time.Sleep(retryDelay)
			fmt.Println("Reintentando conexión...")
		}
	}

	return nil, fmt.Errorf("❌ No se pudo conectar a la base de datos después de %d intentos: %w", retryAttempts, err) // Envolver el error final
}