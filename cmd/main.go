package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/routes"
	"github.com/Alvinferdeveloper/summa-backend/routes/admin"
	"github.com/Alvinferdeveloper/summa-backend/services"
	"github.com/Alvinferdeveloper/summa-backend/websocket"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	config.ConnectDB()
	config.MigrateDatabase()

	services.InitS3Uploader()

	if err := services.InitCacheService(); err != nil {
		log.Fatalf("Failed to initialize cache service: %v", err)
	}

	if err := services.InitEmailService(); err != nil {
		log.Fatalf("Failed to initialize email service: %v", err)
	}

	if err := services.InitRabbitMQService(); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ service: %v", err)
	}

	hub := websocket.NewHub()
	go hub.Run()

	services.GlobalNotificationService = services.NewNotificationService(hub)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/api/v1")
	{
		routes.SetupAuthRoutes(v1)
		routes.SetupProfileRoutes(v1)
		routes.SetupEmployerAuthRoutes(v1)
		routes.SetupJobPostRoutes(v1)
		routes.SetupEmployerRoutes(v1)
		routes.SetupExperienceRoutes(v1)
		routes.SetupJobApplicationRoutes(v1)
		routes.SetupUniversityRoutes(v1)
		routes.SetupUniversitySuggestionRoutes(v1)
		routes.SetupUploadRoutes(v1)
		routes.SetupCategoryRoutes(v1)
		routes.SetupAccessibleInfrastructureRoutes(v1)
		routes.SetupInclusiveProgramRoutes(v1)
		routes.SetupContractTypeRoutes(v1)
		routes.SetupExperienceLevelRoutes(v1)
		routes.SetupWorkScheduleRoutes(v1)
		routes.SetupWorkModelRoutes(v1)
		routes.SetupDisabilityTypeRoutes(v1)
		routes.SetupCandidateRoutes(v1)
		routes.SetupSkillRoutes(v1)
		routes.SetupChatRoutes(v1)
		routes.SetupWebSocketRoutes(v1, hub)
		routes.SetupDashboardRoutes(v1)
		routes.SetupCompatibilityRoutes(v1)
		routes.SetupInterviewRoutes(v1)
		routes.SetupNotificationRoutes(v1)
		routes.SetupAdminAuthRoutes(v1)
		admin.SetupDisabilityTypeRoutes(v1)
		admin.SetupAccessibilityNeedRoutes(v1)
		admin.SetupCategoryRoutes(v1)
		admin.SetupContractTypeRoutes(v1) // Añadir rutas de admin para tipos de contrato
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
