package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupEmployerRoutes registra las rutas relacionadas con los empleadores (búsqueda, etc.).
func SetupEmployerRoutes(router *gin.RouterGroup) {
	employer := router.Group("/employers")
	{
		employer.GET("/search", middlewares.AuthMiddleware("job_seeker", "employer"), controllers.SearchEmployers)
	}
}
