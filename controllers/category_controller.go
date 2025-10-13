package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las categorías"})
		return
	}
	var categoriesDTO []*dto.CategoryResponseDTO
	for _, category := range categories {
		categoriesDTO = append(categoriesDTO, dto.ConvertCategoryToDTO(category))
	}
	c.JSON(http.StatusOK, categoriesDTO)
}
