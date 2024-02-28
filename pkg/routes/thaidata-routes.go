package routes

import (
	"api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var ThaidataRoutes = func(router *gin.Engine) {
	api := router.Group(("api/thaidata"))
	{
		api.GET("/provinces", controllers.GetProvinces)
		api.GET("/amphure/province=:provinceId", controllers.GetAmphureByProvince)
		api.GET("/tambon/amphure=:amphureId", controllers.GetTambonByAmphure)
		api.GET("/zipcode/tambon=:tambonId", controllers.GetZipcodeByTambon)

		// api.GET("/amphures", controllers.GetAmphures)
		// api.GET("/tambons", controllers.GetTambons)
		// api.GET("/zipcodes", controllers.GetZipcodes)
		// api.GET("/province", controllers.SelectedProvince)
	}
}
