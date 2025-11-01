package admin

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers/admin"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupContractTypeRoutes(router *gin.RouterGroup) {
	contractTypes := router.Group("/admin/contract-types")
	contractTypes.Use(middlewares.AuthMiddleware("admin"))
	{
		contractTypes.GET("", admin.GetAllContractTypes)
		contractTypes.POST("", admin.CreateContractType)
		contractTypes.PUT("/:id", admin.UpdateContractType)
		contractTypes.DELETE("/:id", admin.DeleteContractType)
	}
}
