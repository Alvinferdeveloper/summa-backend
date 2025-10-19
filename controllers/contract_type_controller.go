package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

// GetAllContractTypes returns all contract types.
func GetAllContractTypes(c *gin.Context) {
	var contractTypes []models.ContractType
	if err := config.DB.Find(&contractTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve contract types"})
		return
	}

	var dtos []dto.ContractTypeDTO
	for _, ct := range contractTypes {
		dtos = append(dtos, dto.ConvertContractTypeToDTO(ct))
	}

	c.JSON(http.StatusOK, dtos)
}
