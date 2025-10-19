package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupContractTypeRoutes(router *gin.RouterGroup) {
	contractTypeRoutes := router.Group("/contract-types")
	{
		contractTypeRoutes.GET("", controllers.GetAllContractTypes)
	}
}
