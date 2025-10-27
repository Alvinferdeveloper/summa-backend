package controllers

import (
	"net/http"

	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetDashboardStats(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	stats, err := services.GetDashboardStats(employerID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las estadísticas"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func GetPipeline(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	pipeline, err := services.GetPipeline(employerID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el pipeline de contratación"})
		return
	}

	c.JSON(http.StatusOK, pipeline)
}

func GetCandidateSkillInsights(c *gin.Context) {
	employerID, _ := c.Get("employer_id")

	skills, err := services.GetCandidateSkillInsights(employerID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los insights de habilidades"})
		return
	}

	c.JSON(http.StatusOK, skills)
}
