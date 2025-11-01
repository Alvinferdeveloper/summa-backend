
package admin

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

type WorkModelRequest struct {
	Name string `json:"name" binding:"required"`
}

func GetAllWorkModels(c *gin.Context) {
	workModels, err := services.GetAllWorkModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los modelos de trabajo"})
		return
	}
	var dtos []dto.WorkModelDTO
	for _, wm := range workModels {
		dtos = append(dtos, dto.ConvertWorkModelToDTO(wm))
	}
	c.JSON(http.StatusOK, dtos)
}

func CreateWorkModel(c *gin.Context) {
	var req WorkModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	workModel, err := services.CreateWorkModel(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el modelo de trabajo"})
		return
	}
	c.JSON(http.StatusCreated, dto.ConvertWorkModelToDTO(*workModel))
}

func UpdateWorkModel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req WorkModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	workModel, err := services.UpdateWorkModel(uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el modelo de trabajo"})
		return
	}
	c.JSON(http.StatusOK, dto.ConvertWorkModelToDTO(*workModel))
}

func DeleteWorkModel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := services.DeleteWorkModel(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el modelo de trabajo"})
		return	}
	c.JSON(http.StatusOK, gin.H{"message": "Modelo de trabajo eliminado"})
}
