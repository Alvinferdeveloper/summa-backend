package admin

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

type DisabilityTypeRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func GetAllDisabilityTypes(c *gin.Context) {
	disabilityTypes, err := services.GetDisabilityTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tipos de discapacidad"})
		return
	}
	var dtos []dto.DisabilityTypeResponse
	for _, dt := range disabilityTypes {
		dtos = append(dtos, dto.ConvertDisabilityTypeToDTO(dt))
	}
	c.JSON(http.StatusOK, dtos)
}

func CreateDisabilityType(c *gin.Context) {
	var req DisabilityTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	disabilityType, err := services.CreateDisabilityType(req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tipo de discapacidad"})
		return
	}
	c.JSON(http.StatusCreated, dto.ConvertDisabilityTypeToDTO(*disabilityType))
}

func UpdateDisabilityType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req DisabilityTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	disabilityType, err := services.UpdateDisabilityType(uint(id), req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el tipo de discapacidad"})
		return
	}
	c.JSON(http.StatusOK, dto.ConvertDisabilityTypeToDTO(*disabilityType))
}

func DeleteDisabilityType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := services.DeleteDisabilityType(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el tipo de discapacidad"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tipo de discapacidad eliminado"})
}
