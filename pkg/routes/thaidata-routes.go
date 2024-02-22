package routes

import (
	"api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var ThaidataRoutes = func(router *gin.Engine) {
	api := router.Group(("api/thaidata"))
	{
		api.GET("/", controllers.GetThaiData)
	}
}
