package routes

import (
	"api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var CheckboxRoutes = func(router *gin.Engine) {
	api := router.Group("api/checkbox")
	{
		api.POST("/", controllers.CreateCheckbox)
		api.GET("/gender", controllers.GetGender)
		api.GET("/prefix", controllers.GetPrefix)
		api.GET("/position", controllers.GetPosition)
		api.GET("/status", controllers.GetStatus)
	}
}
