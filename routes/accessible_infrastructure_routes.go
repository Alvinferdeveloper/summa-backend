package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupAccessibleInfrastructureRoutes sets up the routes for accessible infrastructures.
func SetupAccessibleInfrastructureRoutes(router *gin.RouterGroup) {
	infrastructureRoutes := router.Group("/infrastructures")
	{
		infrastructureRoutes.GET("", controllers.GetAllAccessibleInfrastructures)
	}
}
