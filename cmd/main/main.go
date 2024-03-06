package main

import (
	// "api/pkg/controllers"
	"api/pkg/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Bearer "}, // เพิ่ม "Authorization" ที่นี่
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// // Middleware to verify token
	// r.Use(func(c *gin.Context) {
	// 	tokenString := c.GetHeader("Authorization")
	// 	if tokenString == "" {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	// 		c.Abort()
	// 		return
	// 	}

	// 	// ตรวจสอบ Token ของคุณได้ที่นี่
	// 	// คุณจำเป็นต้อง import jwt และ secretKey โดยตรงที่นี่
	// 	// และทำการตรวจสอบว่า Token ถูกต้องหรือไม่
	// 	// ซึ่งจะมีการใช้ jwt.Parse() และการตรวจสอบข้อผิดพลาดเช่นเดียวกับตัวอย่างที่ได้แสดงไว้แล้วในตัวอย่างของคุณ

	// 	// หลังจากที่ตรวจสอบ Token เรียบร้อยแล้ว คุณสามารถเรียก c.Next() เพื่อให้เดินต่อไปยัง Middleware หรือเส้นทางถัดไปได้
	// })

	routes.AuthRoutes(r)
	routes.BookRoutes(r)
	routes.UserRoutes(r)
	routes.ThaidataRoutes(r)
	routes.CheckboxRoutes(r)

	port := "8080"
	log.Printf("Server is running at http://localhost:%s\n", port)

	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, r))
}
