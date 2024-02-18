package routes

import (
	"api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterUserRoutes = func(router *gin.Engine) {
	api := router.Group("/api/user")
	{
		// api.Use(verifyToken)
		api.POST("/register", controllers.CreateUser)
		api.POST("/login", controllers.Login)
		api.GET("/", controllers.GetUser)
		api.GET("/:userId", controllers.GetUserById)
		// api.GET("/:email", controllers.GetUserByEmail)
		api.PUT("/:userId", controllers.UpdateUser)
		api.DELETE("/:userId", controllers.DeleteUser)
	}

}
