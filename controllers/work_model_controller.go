package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

// GetAllWorkModels returns all work models.
func GetAllWorkModels(c *gin.Context) {
	var workModels []models.WorkModel
	if err := config.DB.Find(&workModels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve work models"})
		return
	}

	var dtos []dto.WorkModelDTO
	for _, wm := range workModels {
		dtos = append(dtos, dto.ConvertWorkModelToDTO(wm))
	}

	c.JSON(http.StatusOK, dtos)
}
