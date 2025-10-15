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

	jobPostDTO := dto.ConvertJobPostToDTO(*jobPost)

	c.JSON(http.StatusCreated, gin.H{"message": "Job post created successfully", "jobPost": jobPostDTO})
}

func GetJobPosts(c *gin.Context) {
	var page, limit int
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	jobPost, total, hasNextPage, err := services.GetJobPosts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var jobPostDTOs []dto.JobPostResponse
	for _, jobPost := range jobPost {
		jobPostDTOs = append(jobPostDTOs, dto.ConvertJobPostToDTO(jobPost))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      jobPostDTOs,
		"total":     total,
		"page":      page,
		"limit":     limit,
		"next_page": hasNextPage,
	})
}

func GetJobPostById(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleo inválido"})
		return
	}

	jobPost, err := services.GetJobPostById(uint(jobID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jobPostDTO := dto.ConvertJobPostToDTO(*jobPost)

	c.JSON(http.StatusOK, jobPostDTO)
}

func GetEmployerJobPosts(c *gin.Context) {
	employerID, exists := c.Get("employer_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jobPost, err := services.GetJobPostsByEmployerID(employerID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobPost)
}
