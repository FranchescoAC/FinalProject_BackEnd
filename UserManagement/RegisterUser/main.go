package main

import (
	"bytes" // ✅ Agregado para solucionar el error de "undefined: bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"userManagement/model" // ✅ Importa el modelo correctamente
	"userManagement/routes"
)

// ✅ Se debe usar model.User en vez de User
func main() {
	// Registrar rutas
	r := routes.RegisterRoutes()

	// Configurar puerto del servidor
	port := ":2000"
	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// ✅ Se asegura que `model.User` se use correctamente
func registerUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Entrada inválida", http.StatusBadRequest)
		return
	}

	// Crear usuario en la base de datos
	createdUser, err := model.CreateUser(user)
	if err != nil {
		http.Error(w, "Error creando usuario", http.StatusInternalServerError)
		return
	}

	// Enviar Webhook con `model.User`
	sendWebhookToTicketManagement(createdUser)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Usuario registrado exitosamente")
}

// ✅ Función para enviar Webhook a TicketManagement
func sendWebhookToTicketManagement(user model.User) {
	webhookURL := "http://host.docker.internal:4002/api/tickets/register-ticket"

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

	if resp.StatusCode == http.StatusOK {
		log.Println("✅ Webhook enviado exitosamente a TicketManagement")
	} else {
		log.Printf("⚠️ Error enviando Webhook. Estado: %s", resp.Status)
	}
}
