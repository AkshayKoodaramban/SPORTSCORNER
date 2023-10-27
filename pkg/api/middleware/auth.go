package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Replace with your actual secret key.
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "client" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	userIDFloat64, ok := claims["id"].(float64) // Assuming "id" is stored as a float64
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access in id"})
		c.Abort()
		return
	}

	userID := int(userIDFloat64)

	// Set both "role" and "user_id" in the Gin context.
	c.Set("role", role)
	c.Set("id", userID)

	c.Next()

}

// func UserAuthMiddleware(c *gin.Context) {
// 	// Step 1: Retrieve the token from the "Authorization" header
// 	tokenString := c.GetHeader("Authorization")

// 	if tokenString == "" {
// 		// Step 2: Handle the case of a missing authorization token
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
// 		c.Abort()
// 		return
// 	}

// 	// Remove the "Bearer" prefix
// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	// Step 3: Parse the token and check for errors
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("tough"), nil
// 	})

// 	// Step 4: Handle token parsing and validation errors
// 	if err != nil || !token.Valid {
// 		fmt.Println(err)
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
// 		c.Abort()
// 		return
// 	}

// 	// Step 5: Extract claims
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
// 		c.Abort()
// 		return
// 	}

// 	// Step 6: Check the user's role
// 	role, ok := claims["role"].(string)
// 	if !ok || role != "client" {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
// 		c.Abort()
// 		return
// 	}

// 	// Step 7: Extract the user's ID
// 	userID, err := helper.ExtractUserIDFromToken(tokenString, "tough")

// 	// Step 8: Set both "role" and "user_id" in the Gin context
// 	c.Set("role", role)
// 	c.Set("id", userID)

// 	// Allow the request to continue
// 	c.Next()
// }

// admin authentication

func AdminAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	c.Set("role", role)

	c.Next()
}
