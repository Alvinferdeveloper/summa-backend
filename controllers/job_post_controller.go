package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alvinferdeveloper/summa-backend/dto"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/gin-gonic/gin"
)

func CreateJobPost(c *gin.Context) {
	employerID, exists := c.Get("employer_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req dto.CreateJobPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPost, err := services.CreateJobPost(&req, employerID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Job post created successfully", "jobPost": jobPost})
}

func GetJobPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	jobPosts, total, err := services.ListJobPosts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch job posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      jobPosts,
		"total":     total,
		"page":      page,
		"limit":     limit,
		"next_page": page*limit < int(total),
	})
}
