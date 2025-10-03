package controllers

import (
	"net/http"
	"strings"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/gin-gonic/gin"
)

func SearchEmployers(c *gin.Context) {
	query := c.Query("q")
	if len(query) < 2 {
		c.JSON(http.StatusOK, []models.Employer{})
		return
	}

	var employers []models.Employer
	if err := config.DB.Where("LOWER(company_name) LIKE ?", "%"+strings.ToLower(query)+"%").Limit(10).Find(&employers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar empleadores"})
		return
	}

	c.JSON(http.StatusOK, employers)
}
