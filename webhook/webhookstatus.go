package webhook

type WStatus string

const (
	StatusPending   WStatus = "pending"
	StatusProcessed WStatus = "processed"
	StatusFailed    WStatus = "failed"
)

type WebhookStatus struct {
	ID     string                 `json:"id"`     // ID of the webhook
	Status WStatus                `json:"status"` // Status of the webhook
	Data   map[string]interface{} // Added Data field to store webhook data
}
