package routes

import (
	"api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var UserRoutes = func(router *gin.Engine) {
	api := router.Group("/api/user")
	api.Use(verifyToken)
	{
		api.GET("/", controllers.GetUser)
		api.GET("/:userId", controllers.GetUserById)
		api.PUT("/:userId", controllers.UpdateUser)
		api.DELETE("/:userId", controllers.DeleteUser)
	}
}
