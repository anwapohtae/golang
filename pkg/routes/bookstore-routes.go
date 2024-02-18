package routes

import (
	"api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {
	api := router.Group("/api/book", verifyToken)
	{
		api.POST("/", controllers.CreateBook)
		api.GET("/", controllers.GetBook)
		api.GET("/:bookId", controllers.GetBookById)
		api.PUT("/:bookId", controllers.UpdateBook)
		api.DELETE("/:bookId", controllers.DeleteBook)
	}
	
}
