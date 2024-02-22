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
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization = Unauthorized request")
		return
	}

	auth := strings.Split(authHeader, " ")
	if len(auth) < 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Token is missing in Authorization header")
		return
	}

	tokenString := auth[1] // ใช้ index 1 เพื่อรับค่า token หลังช่องว่าง

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-256-bit-secret"), nil // ใส่ "your-256-bit-secret" ด้วย secret key ของคุณ
	})
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid or expired token")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token claims")
		return
	}

	emailUser, ok := claims["email"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid email in token claims")
		return
	}

	c.Set("emailUser", emailUser)
	c.Next()
}
