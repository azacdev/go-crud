package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/azacdev/go-crud/initializers"
	"github.com/azacdev/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie from request
	tokenString, err := c.Cookie("Authorisation")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration time
		if exp, ok := claims["exp"].(float64); ok {
			// Convert exp to time.Time and check if token is expired
			expTime := time.Unix(int64(exp), 0)
			if time.Now().After(expTime) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		var user models.User

		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		c.Set("user", user)

		// Continue to the next handler
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
