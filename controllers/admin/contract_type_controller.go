
package admin

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

type ContractTypeRequest struct {
	Name string `json:"name" binding:"required"`
}

func GetAllContractTypes(c *gin.Context) {
	contractTypes, err := services.GetAllContractTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tipos de contrato"})
		return
	}
	var dtos []dto.ContractTypeDTO
	for _, ct := range contractTypes {
		dtos = append(dtos, dto.ConvertContractTypeToDTO(ct))
	}
	c.JSON(http.StatusOK, dtos)
}

func CreateContractType(c *gin.Context) {
	var req ContractTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractType, err := services.CreateContractType(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tipo de contrato"})
		return
	}
	c.JSON(http.StatusCreated, dto.ConvertContractTypeToDTO(*contractType))
}

func UpdateContractType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var req ContractTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractType, err := services.UpdateContractType(uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el tipo de contrato"})
		return
	}
	c.JSON(http.StatusOK, dto.ConvertContractTypeToDTO(*contractType))
}

func DeleteContractType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := services.DeleteContractType(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el tipo de contrato"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tipo de contrato eliminado"})
}
