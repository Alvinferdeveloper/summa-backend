
package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAdminAuthRoutes(router *gin.RouterGroup) {
	adminAuth := router.Group("/admin")
	{
		adminAuth.POST("/login", controllers.AdminLogin)
	}
}
