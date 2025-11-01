
package admin

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

type AccessibilityNeedRequest struct {
	Name     string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
}

func GetAllAccessibilityNeeds(c *gin.Context) {
	accessibilityNeeds, err := services.GetAllAccessibilityNeeds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las necesidades de accesibilidad"})
		return
	}
	var dtos []dto.AccessibilityNeedResponse
	for _, an := range accessibilityNeeds {
		dtos = append(dtos, dto.ConvertAccessibilityNeedToDTO(an))
	}
	c.JSON(http.StatusOK, dtos)
}

func CreateAccessibilityNeed(c *gin.Context) {
	var req AccessibilityNeedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessibilityNeed, err := services.CreateAccessibilityNeed(req.Name, req.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la necesidad de accesibilidad"})
		return
	}
	c.JSON(http.StatusCreated, dto.ConvertAccessibilityNeedToDTO(*accessibilityNeed))
}

func UpdateAccessibilityNeed(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req AccessibilityNeedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessibilityNeed, err := services.UpdateAccessibilityNeed(uint(id), req.Name, req.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la necesidad de accesibilidad"})
		return
	}
	c.JSON(http.StatusOK, dto.ConvertAccessibilityNeedToDTO(*accessibilityNeed))
}

func DeleteAccessibilityNeed(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := services.DeleteAccessibilityNeed(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la necesidad de accesibilidad"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Necesidad de accesibilidad eliminada"})
}
