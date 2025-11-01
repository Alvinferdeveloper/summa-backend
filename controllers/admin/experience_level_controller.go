
package admin

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

type ExperienceLevelRequest struct {
	Name string `json:"name" binding:"required"`
}

func GetAllExperienceLevels(c *gin.Context) {
	experienceLevels, err := services.GetAllExperienceLevels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los niveles de experiencia"})
		return
	}
	var dtos []dto.ExperienceLevelDTO
	for _, el := range experienceLevels {
		dtos = append(dtos, dto.ConvertExperienceLevelToDTO(el))
	}
	c.JSON(http.StatusOK, dtos)
}

func CreateExperienceLevel(c *gin.Context) {
	var req ExperienceLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	experienceLevel, err := services.CreateExperienceLevel(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el nivel de experiencia"})
		return
	}
	c.JSON(http.StatusCreated, dto.ConvertExperienceLevelToDTO(*experienceLevel))
}

func UpdateExperienceLevel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req ExperienceLevelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	experienceLevel, err := services.UpdateExperienceLevel(uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el nivel de experiencia"})
		return
	}
	c.JSON(http.StatusOK, dto.ConvertExperienceLevelToDTO(*experienceLevel))
}

func DeleteExperienceLevel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := services.DeleteExperienceLevel(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el nivel de experiencia"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Nivel de experiencia eliminado"})
}
