package routes

import (
	"api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var BookRoutes = func(router *gin.Engine) {
    // สร้างกลุ่ม API สำหรับหนังสือและใช้ Middleware สำหรับตรวจสอบ Token
    api := router.Group("/api/book")
    api.Use(verifyToken) // ใช้ Middleware ตรวจสอบ Token กับทุกเส้นทางในกลุ่มนี้
    {
        // กำหนดเส้นทาง API สำหรับการสร้างหนังสือและการดึงข้อมูลอื่น ๆ
        api.POST("/", controllers.CreateBook)
        api.GET("/", controllers.GetBook)
        api.GET("/:bookId", controllers.GetBookById)
        api.PUT("/:bookId", controllers.UpdateBook)
        api.DELETE("/:bookId", controllers.DeleteBook)
    }
}
