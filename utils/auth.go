package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetParticipantFromContext(c *gin.Context) (participantID uuid.UUID, participantType string, err error) {
	if userID, exists := c.Get("user_id"); exists {
		if id, ok := userID.(uuid.UUID); ok {
			return id, "user", nil
		}
	}
	if employerID, exists := c.Get("employer_id"); exists {
		if id, ok := employerID.(uuid.UUID); ok {
			return id, "employer", nil
		}
	}
	return uuid.Nil, "", fmt.Errorf("user not authenticated")
}
