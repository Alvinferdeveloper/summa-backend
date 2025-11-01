package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// AuthMiddleware validates the JWT and sets the user_id in the context.
func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			tokenQuery := c.Query("auth")
			if tokenQuery == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header or query param required"})
				return
			}
			authHeader = fmt.Sprintf("Bearer %s", tokenQuery)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		tokenString := parts[1]
		jwtSecret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			role, ok := claims["role"].(string)
			if !ok {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid role in token"})
				return
			}

			isAllowed := false
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					isAllowed = true
					break
				}
			}
			if !isAllowed {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
				return
			}
			println(role)
			switch role {
			case "job_seeker":
				userID, ok := claims["user_id"].(string)
				if !ok {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id in token"})
					return
				}
				uid, err := uuid.Parse(userID)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id in token"})
					return
				}
				c.Set("user_id", uid)
			case "employer":
				employerID, ok := claims["employer_id"].(string)
				if !ok {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid employer_id in token"})
					return
				}
				uid, err := uuid.Parse(employerID)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid employer_id in token"})
					return
				}
				c.Set("employer_id", uid)
			case "admin":
				adminID, ok := claims["admin_id"].(string)
				if !ok {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin_id in token"})
					return
				}
				uid, err := uuid.Parse(adminID)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin_id in token"})
					return
				}
				c.Set("admin_id", uid)
			default:
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid or unsupported role"})
				return
			}
			c.Set("user_type", role)

			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
	}
}
