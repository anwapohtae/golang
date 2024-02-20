package routes

import (
	"api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterUserRoutes = func(router *gin.Engine) {
	api := router.Group("/api/user")
	{
		api.POST("/register", controllers.CreateUser)
		api.POST("/login", controllers.Login)
		api.GET("/", controllers.GetUser, verifyToken)
		api.GET("/:userId", controllers.GetUserById)
		// api.GET("/:email", controllers.GetUserByEmail)
		api.PUT("/:userId", controllers.UpdateUser, verifyToken)
		api.DELETE("/:userId", controllers.DeleteUser, verifyToken)
	}

}
