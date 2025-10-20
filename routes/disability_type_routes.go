package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDisabilityTypeRoutes(router *gin.RouterGroup) {
	disabilityTypeRoutes := router.Group("/disability-types")
	{
		disabilityTypeRoutes.GET("", controllers.GetAllDisabilityTypes)
	}
}
