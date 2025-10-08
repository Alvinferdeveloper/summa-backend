package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

func SearchEmployers(c *gin.Context) {
	query := c.Query("q")
	if len(query) < 2 {
		c.JSON(http.StatusOK, []dto.EmployerResponseDTO{})
		return
	}

	employers, err := services.SearchEmployers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar empleadores"})
		return
	}

	var employerDTOs []dto.EmployerResponseDTO
	for _, employer := range employers {
		employerDTOs = append(employerDTOs, dto.EmployerResponseDTO{
			ID:          employer.ID,
			CompanyName: employer.CompanyName,
			LogoURL:     employer.LogoURL,
			Industry:    employer.Industry,
		})
	}

	c.JSON(http.StatusOK, employerDTOs)
}
