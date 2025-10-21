package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

func GetCandidates(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "12"))

	filters := make(map[string]string)
	filters["country"] = c.Query("country")
	filters["disability_type_id"] = c.Query("disability_type_id")

	profiles, total, err := services.GetCandidates(page, limit, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var candidateDTOs []dto.CandidateResponseDTO
	for _, profile := range profiles {
		candidateDTOs = append(candidateDTOs, dto.ConvertProfileToCandidateDTO(profile))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  candidateDTOs,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
