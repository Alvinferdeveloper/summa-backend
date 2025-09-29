
package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes registers the authentication-related routes.
func SetupAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/google/callback", controllers.OAuthCallbackHandler)
	}
}
