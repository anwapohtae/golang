package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// // Middleware function for token verification
// func verifyToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")

// 		// Check if the token is valid
// 		if token != "your_secret_token" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}

// 		// Token is valid, continue to the next handler
// 		c.Next()
// 	}
// }

func verifyToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized request")
		return
	}

	tokenString := strings.Split(authHeader, " ")[1]
	if tokenString == "null" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized request")
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil // replace "secretKey" with your secret key
	})
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized request")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized request")
		return
	}

	userId, ok := claims["subject"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized request")
		return
	}

	c.Set("userId", userId)
	c.Next()
}
