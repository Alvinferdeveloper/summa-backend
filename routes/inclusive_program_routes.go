package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupInclusiveProgramRoutes sets up the routes for inclusive programs.
func SetupInclusiveProgramRoutes(router *gin.RouterGroup) {
	programRoutes := router.Group("/programs")
	{
		programRoutes.GET("", controllers.GetAllInclusivePrograms)
	}
}
