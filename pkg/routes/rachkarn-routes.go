package routes

import (
	"api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RachkarnRoutes = func(router *gin.Engine) {
	api := router.Group("/api/rachkarn")
	{
		api.POST("/", controllers.CreateRachkarn)
		api.GET("/", controllers.GetRachkarn)
		api.GET("/:userid", controllers.GetRachkarnByUserId)
	}
}