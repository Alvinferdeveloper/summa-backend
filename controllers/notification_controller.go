package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetNotifications(c *gin.Context) {
	userTypeVal, ok := c.Get("user_type")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}

	userType, ok := userTypeVal.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}

	switch userType {
	case "job_seeker":
		userIDVal, ok := c.Get("user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}
		userID, ok := userIDVal.(uuid.UUID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}
		notifications, err := services.GetNotificationsForUser(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener notificaciones"})
			return
		}
		c.JSON(http.StatusOK, notifications)

	case "employer":
		employerIDVal, ok := c.Get("employer_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}
		employerID, ok := employerIDVal.(uuid.UUID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}
		notifications, err := services.GetNotificationsForEmployer(employerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener notificaciones"})
			return
		}

		notificationResponses := make([]dto.NotificationResponse, 0)
		for _, notification := range notifications {
			notificationResponses = append(notificationResponses, dto.ConvertNotificationToNotificationResponse(notification))
		}
		c.JSON(http.StatusOK, notificationResponses)

	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
	}
}

func MarkNotificationsAsRead(c *gin.Context) {
	userTypeVal, ok := c.Get("user_type")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}
	userType, ok := userTypeVal.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}

	var req struct {
		IDs []uint `json:"ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDs inválidos"})
		return
	}

	var err error
	switch userType {
	case "job_seeker":
		userIDVal, ok := c.Get("user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}
		uuidUserID, ok := userIDVal.(uuid.UUID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}
		err = services.MarkNotificationsAsRead(req.IDs, &uuidUserID, nil)

	case "employer":
		employerIDVal, ok := c.Get("employer_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}
		uuidEmployerID, ok := employerIDVal.(uuid.UUID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}
		err = services.MarkNotificationsAsRead(req.IDs, nil, &uuidEmployerID)

	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al marcar notificaciones como leídas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notificaciones marcadas como leídas"})
}
