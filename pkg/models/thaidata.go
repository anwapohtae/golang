package models

import (
	"api/pkg/config"

	"gorm.io/gorm"
)

var dtLocation *gorm.DB

type Location struct {
	gorm.Model
	ProvinceID   string `json:"province_id"`
	ProvinceNameth string `json:"province_nameth"`
	AmphureID    string `json:"amphure_id"`
	AmphureNameth  string `json:"amphure_nameth"`
	TambonID     string `json:"tambon_id"`
	TambonNameth   string `json:"tambon_nameth"`
	Zipcode      string `json:"zipcode"`
}
 
func init() {
	config.Connect()
	dtLocation = config.GetDB()
	dtLocation.AutoMigrate(&Location{})
}

func GetAllThaiData() []Location {
	var ThaiDatas []Location
	dtLocation.Find(&ThaiDatas)
	return ThaiDatas
}
