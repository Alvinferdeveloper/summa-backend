package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllInclusivePrograms(c *gin.Context) {
	var programs []models.InclusiveProgram
	if err := config.DB.Find(&programs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inclusive programs"})
		return
	}

	var dtos []dto.InclusiveProgramDTO
	for _, prog := range programs {
		dtos = append(dtos, dto.ConvertInclusiveProgramToDTO(prog))
	}

	c.JSON(http.StatusOK, dtos)
}
