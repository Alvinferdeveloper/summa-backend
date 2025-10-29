package dto

import (
	"time"

	"github.com/Alvinferdeveloper/summa-backend/models"
)

type NotificationResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	Link      string    `json:"link"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

func ConvertNotificationToNotificationResponse(notification models.Notification) NotificationResponse {
	return NotificationResponse{
		ID:        notification.ID,
		Message:   notification.Message,
		Link:      notification.Link,
		IsRead:    notification.IsRead,
		CreatedAt: notification.CreatedAt,
	}
}
