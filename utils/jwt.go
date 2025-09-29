package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a new JWT for a given user ID and their onboarding status.
func GenerateJWT(userID uint, onboardingCompleted bool) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", fmt.Errorf("missing environment variable %q", "JWT_SECRET")
	}

	claims := jwt.MapClaims{
		"user_id":              userID,
		"role":                 "job_seeker",
		"onboarding_completed": onboardingCompleted,
		"exp":                  time.Now().Add(time.Hour * 72).Unix(),
		"iat":                  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
