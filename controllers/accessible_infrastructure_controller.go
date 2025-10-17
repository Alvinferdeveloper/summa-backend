package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllAccessibleInfrastructures(c *gin.Context) {
	var infrastructures []models.AccessibleInfrastructure
	if err := config.DB.Find(&infrastructures).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accessible infrastructures"})
		return
	}

	var dtos []dto.AccessibleInfrastructureDTO
	for _, infra := range infrastructures {
		dtos = append(dtos, dto.ConvertAccessibleInfrastructureToDTO(infra))
	}

	c.JSON(http.StatusOK, dtos)
}
