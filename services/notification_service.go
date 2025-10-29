package services

import (
	"encoding/json"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/google/uuid"
)

type Broadcaster interface {
	BroadcastToUser(userID string, payload []byte)
}

type NotificationService struct {
	Broadcaster Broadcaster
}

func NewNotificationService(broadcaster Broadcaster) *NotificationService {
	return &NotificationService{Broadcaster: broadcaster}
}

func (s *NotificationService) CreateNotification(notification models.Notification) error {
	if err := config.DB.Create(&notification).Error; err != nil {
		return err
	}

	var targetID uuid.UUID
	if notification.UserID != nil {
		targetID = *notification.UserID
	} else if notification.EmployerID != nil {
		targetID = *notification.EmployerID
	}

	if targetID != uuid.Nil && s.Broadcaster != nil {
		notificationBytes, err := json.Marshal(notification)
		if err != nil {
			return err
		}
		s.Broadcaster.BroadcastToUser(targetID.String(), notificationBytes)
	}

	return nil
}

func GetNotificationsForUser(userID uuid.UUID) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := config.DB.Where("user_id = ?", userID).Order("created_at desc").Limit(20).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func GetNotificationsForEmployer(employerID uuid.UUID) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := config.DB.Where("employer_id = ?", employerID).Order("created_at desc").Limit(20).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

func MarkNotificationsAsRead(ids []uint, userID *uuid.UUID, employerID *uuid.UUID) error {
	query := config.DB.Model(&models.Notification{}).Where("id IN ?", ids)

	if userID != nil {
		query = query.Where("user_id = ?", userID)
	} else if employerID != nil {
		query = query.Where("employer_id = ?", employerID)
	} else {
		return nil
	}

	return query.Update("is_read", true).Error
}
