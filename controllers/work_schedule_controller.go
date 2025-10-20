package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

// GetAllWorkSchedules returns all work schedules.
func GetAllWorkSchedules(c *gin.Context) {
	var workSchedules []models.WorkSchedule
	if err := config.DB.Find(&workSchedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve work schedules"})
		return
	}

	var dtos []dto.WorkScheduleDTO
	for _, ws := range workSchedules {
		dtos = append(dtos, dto.ConvertWorkScheduleToDTO(ws))
	}

	c.JSON(http.StatusOK, dtos)
}
