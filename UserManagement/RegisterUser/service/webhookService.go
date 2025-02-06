package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"userManagement/model"
)

// Enviar el webhook a TicketManagement
func SendWebhookToTicketManagement(user model.User) {
	webhookURL := "http://host.docker.internal:4002/api/tickets/register-ticket"

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

	// Crear y enviar la solicitud HTTP POST
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
