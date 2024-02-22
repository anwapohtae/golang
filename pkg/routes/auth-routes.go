package routes

import (
	"api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var AuthRoutes = func(router *gin.Engine) {
	api := router.Group("/api/auth")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

	}
}