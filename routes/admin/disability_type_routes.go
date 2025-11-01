package admin

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers/admin"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupDisabilityTypeRoutes(router *gin.RouterGroup) {
	disabilityTypes := router.Group("/admin/disability-types")
	disabilityTypes.Use(middlewares.AuthMiddleware("admin"))
	{
		disabilityTypes.GET("", admin.GetAllDisabilityTypes)
		disabilityTypes.POST("", admin.CreateDisabilityType)
		disabilityTypes.PUT("/:id", admin.UpdateDisabilityType)
		disabilityTypes.DELETE("/:id", admin.DeleteDisabilityType)
	}
}
