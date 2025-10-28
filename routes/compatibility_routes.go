package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupCompatibilityRoutes(router *gin.RouterGroup) {
	router.GET("/jobs/:id/compatibility", middlewares.AuthMiddleware("job_seeker"), controllers.GetJobCompatibility)
}
