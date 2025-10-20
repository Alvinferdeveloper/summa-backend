package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

func GetAllDisabilityTypes(c *gin.Context) {
	disabilityTypes, err := services.GetDisabilityTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch disability types"})
		return
	}

	var dtos []dto.DisabilityTypeResponse
	for _, dt := range disabilityTypes {
		dtos = append(dtos, dto.ConvertDisabilityTypeToDTO(dt))
	}

	c.JSON(http.StatusOK, dtos)
}
