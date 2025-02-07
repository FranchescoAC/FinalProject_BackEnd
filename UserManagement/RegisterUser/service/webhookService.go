package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"userManagement/model"
)

// SendWebhookToNotificationManagement sends a webhook to NotificationManagement when a new user is registered
func SendWebhookToNotificationManagement(user model.User) {
	// Change this URL to NotificationManagement's webhook endpoint
	//webhookURL := "http://host.docker.internal:5000/api/notifications/new-user"
	webhookURL := "http://localhost:5000/api/notifications/new-user"

	// Data to be sent in the webhook
	webhookData := map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	}

	// Convert the payload to JSON
	jsonData, err := json.Marshal(webhookData)
	if err != nil {
		log.Println("❌ Error marshaling user data to JSON:", err)
		return
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("❌ Error creating the HTTP request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("❌ Error sending the webhook:", err)
		return
	}
	defer resp.Body.Close()

	// Log the response
	if resp.StatusCode == http.StatusOK {
		log.Println("✅ Webhook successfully sent to NotificationManagement")
	} else {
		log.Printf("⚠️ Failed to send webhook. Status: %s", resp.Status)
	}
}
