
package routes

import (
	"github.com/Alvinferdeveloper/summa-backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetupEmployerAuthRoutes registers the employer authentication-related routes.
func SetupEmployerAuthRoutes(router *gin.RouterGroup) {
	employerAuth := router.Group("/employer")
	{
		employerAuth.POST("/register", controllers.RegisterEmployer)
		employerAuth.POST("/login", controllers.LoginEmployer)
	}
}
