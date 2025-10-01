package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/routes"
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
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
