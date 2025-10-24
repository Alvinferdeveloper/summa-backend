package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func GetMyEmployerProfile(c *gin.Context) {
	employerID, exists := c.Get("employer_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	employer, err := services.FindEmployerByID(employerID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Perfil de empleador no encontrado"})
		return
	}
	employerDTO := dto.ConvertEmployerToDTO(*employer)

	c.JSON(http.StatusOK, employerDTO)
}

func UpdateMyEmployerProfile(c *gin.Context) {
	employerID, exists := c.Get("employer_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req dto.UpdateEmployerProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employer, err := services.UpdateEmployerProfile(employerID.(uuid.UUID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el perfil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfil actualizado exitosamente", "employer": employer})
}
