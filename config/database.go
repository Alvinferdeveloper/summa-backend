package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Alvinferdeveloper/summa-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connection successfully opened.")
}

// MigrateDatabase runs the auto-migration for all models.
func MigrateDatabase() {
	fmt.Println("Running database migrations...")
	err := DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Skill{},
		&models.Employer{},
		&models.Experience{},
		&models.University{},
		&models.ProfileEducation{},
		&models.UniversitySuggestion{},
		&models.NewEmployer{},
		&models.ProfileSkill{},
		&models.DisabilityType{},
		&models.AccessibilityNeed{},
		&models.JobPost{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migration completed.")
}
