package webhook

import (
	"fmt"
	"log"
	"sync"
)

type WebhookHandler struct {
	webhookStatusMap  map[string]WebhookStatus // Map to store webhook statuses
	webhookStatusLock sync.Mutex               // Mutex to protect the map
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{
		webhookStatusMap: make(map[string]WebhookStatus),
	}
}

func (wh *WebhookHandler) GetWebhookStatus(id string) (WebhookStatus, bool) {
	wh.webhookStatusLock.Lock()
	defer wh.webhookStatusLock.Unlock()
	status, exists := wh.webhookStatusMap[id]
	return status, exists
}

func (wh *WebhookHandler) ProcessWehook(id string, data map[string]interface{}) error {
	wh.webhookStatusLock.Lock()
	defer wh.webhookStatusLock.Unlock()
	// check if the webhook is already processed
	if _, exists := wh.webhookStatusMap[id]; exists {
		return fmt.Errorf("Already processed")
	}
	// Simulate webhook processing
	wh.webhookStatusMap[id] = WebhookStatus{
		ID:     id,
		Status: StatusProcessed,
		Data:   data,
	}
	log.Default().Println("Webhook processed: ", id)
	return nil
}
