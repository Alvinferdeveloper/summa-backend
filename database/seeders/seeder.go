package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Alvinferdeveloper/summa-backend/config"
	"github.com/Alvinferdeveloper/summa-backend/models"
	"github.com/Alvinferdeveloper/summa-backend/utils"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load() // Load .env file from current directory or parents
	if err != nil {
		log.Println("Error loading .env file, using environment variables: ", err)
	}

	config.ConnectDB()

	config.MigrateDatabase()

	err = RunSeeder(config.DB)
	if err != nil {
		log.Fatalf("Seeder failed: %v", err)
	}
	fmt.Println("Database seeding completed successfully!")
}

func RunSeeder(db *gorm.DB) error {
	fmt.Println("Starting database seeding...")

	// Clear existing data (optional, for clean runs)
	// This should be used with caution in production!
	// db.Exec("DELETE FROM profile_skills")
	// db.Exec("DELETE FROM accessibility_needs")
	// db.Exec("DELETE FROM disability_types")
	// db.Exec("DELETE FROM skills")
	// db.Exec("DELETE FROM profile_educations")
	// db.Exec("DELETE FROM experiences")
	// db.Exec("DELETE FROM job_posts")
	// db.Exec("DELETE FROM employers")
	// db.Exec("DELETE FROM profiles")
	// db.Exec("DELETE FROM users")
	// db.Exec("DELETE FROM universities")
	// db.Exec("DELETE FROM university_suggestions")
	// db.Exec("DELETE FROM new_employers")

	// Seed basic types first
	disabilityTypes, err := seedDisabilityTypes(db)
	if err != nil {
		return fmt.Errorf("failed to seed disability types: %w", err)
	}
	accessibilityNeeds, err := seedAccessibilityNeeds(db)
	if err != nil {
		return fmt.Errorf("failed to seed accessibility needs: %w", err)
	}
	universities, err := seedUniversities(db)
	if err != nil {
		return fmt.Errorf("failed to seed universities: %w", err)
	}
	skills, err := seedSkills(db)
	if err != nil {
		return fmt.Errorf("failed to seed skills: %w", err)
	}

	// Seed users and employers
	users, err := seedUsers(db)
	if err != nil {
		return fmt.Errorf("failed to seed users: %w", err)
	}
	employers, err := seedEmployers(db)
	if err != nil {
		return fmt.Errorf("failed to seed employers: %w", err)
	}

	// Seed dependent models
	err = seedProfiles(db, users, disabilityTypes, accessibilityNeeds, skills)
	if err != nil {
		return fmt.Errorf("failed to seed profiles: %w", err)
	}
	err = seedJobPosts(db, employers, accessibilityNeeds)
	if err != nil {
		return fmt.Errorf("failed to seed job posts: %w", err)
	}
	err = seedExperiences(db, users, employers)
	if err != nil {
		return fmt.Errorf("failed to seed experiences: %w", err)
	}
	err = seedProfileEducations(db, users, universities)
	if err != nil {
		return fmt.Errorf("failed to seed profile educations: %w", err)
	}

	return nil
}

// --- Seeding Functions for Each Model ---

func seedDisabilityTypes(db *gorm.DB) ([]models.DisabilityType, error) {
	fmt.Println("Seeding Disability Types...")
	disabilityTypes := []models.DisabilityType{
		{Name: "Visual", Description: "Discapacidad relacionada con la vista."},
		{Name: "Auditiva", Description: "Discapacidad relacionada con el oído."},
		{Name: "Motriz", Description: "Discapacidad que afecta la movilidad."},
		{Name: "Cognitiva", Description: "Discapacidad que afecta el procesamiento de información."},
		{Name: "Del Habla", Description: "Discapacidad que afecta la comunicación verbal."},
	}
	for i := range disabilityTypes {
		if err := db.FirstOrCreate(&disabilityTypes[i], models.DisabilityType{Name: disabilityTypes[i].Name}).Error; err != nil {
			return nil, err
		}
	}
	return disabilityTypes, nil
}

func seedAccessibilityNeeds(db *gorm.DB) ([]models.AccessibilityNeed, error) {
	fmt.Println("Seeding Accessibility Needs...")
	accessibilityNeeds := []models.AccessibilityNeed{
		{Name: "Lector de pantalla", Category: "Herramientas Digitales"},
		{Name: "Software de magnificación", Category: "Herramientas Digitales"},
		{Name: "Teclado ergonómico", Category: "Entorno Físico"},
		{Name: "Rampas de acceso", Category: "Entorno Físico"},
		{Name: "Horarios flexibles", Category: "Flexibilidad Laboral"},
		{Name: "Trabajo remoto", Category: "Flexibilidad Laboral"},
		{Name: "Intérprete de lengua de señas", Category: "Apoyo Humano"},
		{Name: "Baños accesibles", Category: "Entorno Físico"},
	}
	for i := range accessibilityNeeds {
		if err := db.FirstOrCreate(&accessibilityNeeds[i], models.AccessibilityNeed{Name: accessibilityNeeds[i].Name}).Error; err != nil {
			return nil, err
		}
	}
	return accessibilityNeeds, nil
}

func seedUniversities(db *gorm.DB) ([]models.University, error) {
	fmt.Println("Seeding Universities...")
	universities := []models.University{
		{Name: "Universidad Nacional Autónoma de México", Country: "México", Website: "https://www.unam.mx"},
		{Name: "Universidad de Buenos Aires", Country: "Argentina", Website: "https://www.uba.ar"},
		{Name: "Universidad Complutense de Madrid", Country: "España", Website: "https://www.ucm.es"},
	}
	for i := range universities {
		if err := db.FirstOrCreate(&universities[i], models.University{Name: universities[i].Name}).Error; err != nil {
			return nil, err
		}
	}
	return universities, nil
}

func seedSkills(db *gorm.DB) ([]models.Skill, error) {
	fmt.Println("Seeding Skills...")
	skills := []models.Skill{
		{Name: "Go"},
		{Name: "React"},
		{Name: "JavaScript"},
		{Name: "TypeScript"},
		{Name: "PostgreSQL"},
		{Name: "Docker"},
		{Name: "Kubernetes"},
		{Name: "Comunicación"},
		{Name: "Trabajo en Equipo"},
	}
	for i := range skills {
		if err := db.FirstOrCreate(&skills[i], models.Skill{Name: skills[i].Name}).Error; err != nil {
			return nil, err
		}
	}
	return skills, nil
}

func seedUsers(db *gorm.DB) ([]models.User, error) {
	fmt.Println("Seeding Users...")
	users := []models.User{
		{Email: "jobseeker1@example.com", Provider: "google", ProviderID: "google_id_1"},
		{Email: "jobseeker2@example.com", Provider: "google", ProviderID: "google_id_2"},
	}
	for i := range users {
		if err := db.FirstOrCreate(&users[i], models.User{Email: users[i].Email}).Error; err != nil {
			return nil, err
		}
	}
	return users, nil
}

func seedEmployers(db *gorm.DB) ([]models.Employer, error) {
	fmt.Println("Seeding Employers...")
	hashedPassword, _ := utils.HashPassword("Password123!")
	employers := []models.Employer{
		{CompanyName: "Tech Solutions Inc.", Email: "employer1@example.com", Password: hashedPassword, Role: "employer", Website: "https://techsolutions.com"},
		{CompanyName: "Innovate Corp.", Email: "employer2@example.com", Password: hashedPassword, Role: "employer", Website: "https://innovate.com"},
	}
	for i := range employers {
		if err := db.FirstOrCreate(&employers[i], models.Employer{Email: employers[i].Email}).Error; err != nil {
			return nil, err
		}
	}
	return employers, nil
}

func seedProfiles(db *gorm.DB, users []models.User, dTypes []models.DisabilityType, aNeeds []models.AccessibilityNeed, skills []models.Skill) error {
	fmt.Println("Seeding Profiles...")
	for i, user := range users {
		profile := models.Profile{
			UserID:              user.ID,
			FirstName:           "Nombre" + fmt.Sprintf("%d", i+1),
			LastName:            "Apellido" + fmt.Sprintf("%d", i+1),
			OnboardingCompleted: true,
			PhoneNumber:         fmt.Sprintf("555-123-%03d", i),
			City:                "Ciudad",
			Country:             "País",
			Description:         "Soy un buscador de empleo entusiasta con experiencia en...",
			LinkedIn:            fmt.Sprintf("https://linkedin.com/in/jobseeker%d", i+1),
			ResumeURL:           fmt.Sprintf("https://example.com/resume%d.pdf", i+1),
		}
		// Ensure profile is created or found
		if err := db.FirstOrCreate(&profile, models.Profile{UserID: user.ID}).Error; err != nil {
			return err
		}

		// Associate some disability types and accessibility needs
		if len(dTypes) > 0 {
			db.Model(&profile).Association("DisabilityTypes").Replace(dTypes[i%len(dTypes)])
		}
		if len(aNeeds) > 0 {
			db.Model(&profile).Association("AccessibilityNeeds").Replace(aNeeds[i%len(aNeeds)])
		}
		// Associate some skills
		if len(skills) > 0 {
			db.Model(&profile).Association("Skills").Replace(skills[i%len(skills)])
		}
	}
	return nil
}

func seedJobPosts(db *gorm.DB, employers []models.Employer, aNeeds []models.AccessibilityNeed) error {
	fmt.Println("Seeding Job Posts...")
	for i, employer := range employers {
		features := []string{}
		if len(aNeeds) > 0 {
			features = append(features, aNeeds[i%len(aNeeds)].Name)
		}
		featuresJSON, _ := json.Marshal(features)

		jobPost := models.JobPost{
			EmployerID:            employer.ID,
			Title:                 fmt.Sprintf("Desarrollador Go Senior %d", i+1),
			Location:              "Remoto",
			WorkModel:             "Remoto",
			ContractType:          "Tiempo Completo",
			Description:           "Buscamos un desarrollador Go experimentado para unirse a nuestro equipo.",
			Responsibilities:      "Diseñar, construir y mantener código Go eficiente, reutilizable y fiable.",
			Requirements:          "Más de 5 años de experiencia en Go. Conocimiento de PostgreSQL.",
			AccessibilityFeatures: string(featuresJSON),
		}
		if err := db.FirstOrCreate(&jobPost, models.JobPost{EmployerID: employer.ID, Title: jobPost.Title}).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedExperiences(db *gorm.DB, users []models.User, employers []models.Employer) error {
	fmt.Println("Seeding Experiences...")
	for i, user := range users {
		var profile models.Profile
		if err := db.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
			return err
		}

		startDate := time.Now().AddDate(-5, 0, 0)
		endDate := time.Now().AddDate(-2, 0, 0)

		exp := models.Experience{
			ProfileID:   profile.ID,
			EmployerID:  &employers[i%len(employers)].ID,
			JobTitle:    fmt.Sprintf("Ingeniero de Software %d", i+1),
			Description: "Desarrollo de aplicaciones backend con Go y microservicios.",
			StartDate:   startDate,
			EndDate:     &endDate,
		}
		if err := db.FirstOrCreate(&exp, models.Experience{ProfileID: profile.ID, JobTitle: exp.JobTitle}).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedProfileEducations(db *gorm.DB, users []models.User, universities []models.University) error {
	fmt.Println("Seeding Profile Educations...")
	for i, user := range users {
		var profile models.Profile
		if err := db.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
			return err
		}

		startDate := time.Now().AddDate(-10, 0, 0)
		endDate := time.Now().AddDate(-6, 0, 0)

		edu := models.ProfileEducation{
			ProfileID:    profile.ID,
			UniversityID: &universities[i%len(universities)].ID,
			Degree:       "Ingeniería en Sistemas",
			FieldOfStudy: "Desarrollo de Software",
			StartDate:    startDate,
			EndDate:      &endDate,
		}
		if err := db.FirstOrCreate(&edu, models.ProfileEducation{ProfileID: profile.ID, Degree: edu.Degree}).Error; err != nil {
			return err
		}
	}
	return nil
}
