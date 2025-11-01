package admin

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers/admin"
	"github.com/Alvinferdeveloper/summa-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupAccessibilityNeedRoutes(router *gin.RouterGroup) {
	accessibilityNeeds := router.Group("/admin/accessibility-needs")
	accessibilityNeeds.Use(middlewares.AuthMiddleware("admin"))
	{
		accessibilityNeeds.GET("", admin.GetAllAccessibilityNeeds)
		accessibilityNeeds.POST("", admin.CreateAccessibilityNeed)
		accessibilityNeeds.PUT("/:id", admin.UpdateAccessibilityNeed)
		accessibilityNeeds.DELETE("/:id", admin.DeleteAccessibilityNeed)
	}
}
