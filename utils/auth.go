package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetParticipantFromContext(c *gin.Context) (participantID uint, participantType string, err error) {
	if userID, exists := c.Get("user_id"); exists {
		if id, ok := userID.(uint); ok {
			return id, "user", nil
		}
	}
	if employerID, exists := c.Get("employer_id"); exists {
		if id, ok := employerID.(uint); ok {
			return id, "employer", nil
		}
	}
	return 0, "", fmt.Errorf("user not authenticated")
}
