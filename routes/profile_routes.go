package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupProfileRoutes(router *gin.RouterGroup) {
	router.GET("/accessibility-needs", controllers.GetAccessibilityNeeds)

	profile := router.Group("/profile")
	profile.Use(middlewares.AuthMiddleware("job_seeker"))
	{
		profile.GET("/me", controllers.GetMyProfile)
		profile.PUT("", controllers.CompleteOnboarding)
		profile.PUT("/personal-info", controllers.UpdatePersonalInfo)
		profile.PUT("/contact-info", controllers.UpdateContactInfo)
		profile.PUT("/description", controllers.UpdateDescription)
		profile.PUT("/disability-info", controllers.UpdateDisabilityInfo)
		profile.PUT("/skills", controllers.UpdateSkills)

		// Experience routes
		profile.POST("/experiences", controllers.CreateExperience)
		profile.PUT("/experiences/:id", controllers.UpdateExperience)
		profile.DELETE("/experiences/:id", controllers.DeleteExperience)

		// Education routes
		profile.POST("/educations", controllers.CreateEducation)
		profile.PUT("/educations/:id", controllers.UpdateEducation)
		profile.DELETE("/educations/:id", controllers.DeleteEducation)

		// New Employer Suggestion
		profile.POST("/new-employer-suggestion", controllers.SuggestNewEmployer)
	}

	// Rutas para que los empleadores vean perfiles de candidatos
	applicants := router.Group("/applicants")
	{
		applicants.GET("/:id/profile", middlewares.AuthMiddleware("employer"), controllers.GetCandidateProfileForEmployer)
	}
}
