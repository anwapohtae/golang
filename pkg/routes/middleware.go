package routes

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func verifyToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "authHeader = Unauthorized request")
		return
	}

	tokenString := strings.Split(authHeader, " ")[1]
	// log.Fatal(tokenString)
	if tokenString == "null" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "tokenString = Unauthorized request")
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-256-bit-secret"), nil // replace "secretKey" with your secret key
	})
	// log.Println(token)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "token = Unauthorized request")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	// log.Println(claims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "claims = Unauthorized request")
		return
	}

	emailUser, ok := claims["email"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "emailUser = Unauthorized request")
		return
	}

	c.Set("emailUser", emailUser)
	c.Next()
}
