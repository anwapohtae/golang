package routes

import (
	"api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var BookRoutes = func(router *gin.Engine) {
    api := router.Group("/api/book")
    {
        api.POST("/", controllers.CreateBook)
        api.GET("/", controllers.GetBook)
        api.GET("/:bookId", controllers.GetBookById)
        api.PUT("/:bookId", controllers.UpdateBook)
        api.DELETE("/:bookId", controllers.DeleteBook)
    }
}
