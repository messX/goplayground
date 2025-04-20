package webhook

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var (
	// WebhookHandler is a global instance of the WebhookHandler
	webhookHandler *WebhookHandler
	once           sync.Once
)

func getWebhookHandler() *WebhookHandler {
	once.Do(func() {

		webhookHandler = &WebhookHandler{
			webhookStatusMap: make(map[string]WebhookStatus),
		}
	})
	return webhookHandler
}

func Run() {
	log.Println("Starting webhook server...")
	http.HandleFunc("/webhook/", webhookHandlerFunc)
	http.ListenAndServe(":8080", nil)
}

func webhookHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if !CheckBasicAuth(r, "username", "password") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idempotencyID := r.URL.Path[len("/webhook/"):]
	if idempotencyID == "" {
		http.Error(w, "idempotencyId is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPost:
		handlePost(w, r, idempotencyID)
	case http.MethodGet:
		handleGet(w, r, idempotencyID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request, idempotencyID string) {
	webhookHandler = getWebhookHandler()
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	if err := webhookHandler.ProcessWehook(idempotencyID, payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func handleGet(w http.ResponseWriter, r *http.Request, idempotencyID string) {
	webhookHandler = getWebhookHandler()
	status, exists := webhookHandler.GetWebhookStatus(idempotencyID)
	if !exists {
		http.Error(w, "Webhook not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook status retrieved successfully"))
	log.Default().Println("Webhook status retrieved: ", status)
	return
}
