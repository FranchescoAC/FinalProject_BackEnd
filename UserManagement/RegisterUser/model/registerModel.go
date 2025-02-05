package model

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"userManagement/config"

	_ "github.com/lib/pq"
)

// Definición de la estructura del usuario
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Función para crear un nuevo usuario en la base de datos y enviar Webhook
func CreateUser(user User) (User, error) {
	// Conectamos con la base de datos
	db, err := config.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close() // Cerramos la conexión al final

	// Query SQL para insertar un nuevo usuario
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	var userID int
	err = db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return User{}, err
	}

	// Asignamos el ID generado a la estructura User
	user.ID = userID

	// Llamamos a la función que envía el Webhook a TicketManagement
	go sendWebhookToTicketManagement(user)

	return user, nil
}

// Función para enviar el Webhook a TicketManagement
func sendWebhookToTicketManagement(user User) {
	webhookURL := "http://host.docker.internal:4002/api/tickets/register-ticket" // URL del Webhook

	// Datos a enviar en el Webhook
	webhookData := map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	}

	jsonData, err := json.Marshal(webhookData)
	if err != nil {
		log.Println("❌ Error al convertir usuario a JSON:", err)
		return
	}

	// Crear solicitud HTTP POST
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("❌ Error creando la solicitud HTTP:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("❌ Error enviando Webhook:", err)
		return
	}
	defer resp.Body.Close()

	// Verificar respuesta del Webhook
	if resp.StatusCode == http.StatusOK {
		log.Println("✅ Webhook enviado exitosamente a TicketManagement")
	} else {
		log.Printf("⚠️ Error enviando Webhook. Estado: %s", resp.Status)
	}
}

// Agrega una función que use explícitamente el paquete `database/sql`
func ExampleSQLUsage() *sql.DB {
	db, _ := sql.Open("postgres", "connection_string")
	return db
}
