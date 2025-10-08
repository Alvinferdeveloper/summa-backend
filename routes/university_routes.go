package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupUniversityRoutes(router *gin.RouterGroup) {
	university := router.Group("/universities")
	{
		university.GET("/search", middlewares.AuthMiddleware("job_seeker", "employer"), controllers.SearchUniversities)
	}
}
