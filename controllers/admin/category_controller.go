
package admin

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las categorías"})
		return
	}
	var dtos []dto.CategoryResponseDTO
	for _, cat := range categories {
		dtos = append(dtos, *dto.ConvertCategoryToDTO(cat))
	}
	c.JSON(http.StatusOK, dtos)
}

func CreateCategory(c *gin.Context) {
	var req CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := services.CreateCategory(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la categoría"})
		return
	}
	c.JSON(http.StatusCreated, dto.ConvertCategoryToDTO(*category))
}

func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := services.UpdateCategory(uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la categoría"})
		return
	}
	c.JSON(http.StatusOK, dto.ConvertCategoryToDTO(*category))
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := services.DeleteCategory(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la categoría"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categoría eliminada"})
}
