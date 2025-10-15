package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupEmployerRoutes(router *gin.RouterGroup) {
	employer := router.Group("/employers")
	{
		employer.GET("/search", middlewares.AuthMiddleware("job_seeker", "employer"), controllers.SearchEmployers)
	}

	me := router.Group("/employer/me")
	me.Use(middlewares.AuthMiddleware("employer"))
	{
		me.GET("", controllers.GetMyEmployerProfile)
		me.PUT("", controllers.UpdateMyEmployerProfile)
	}
}
