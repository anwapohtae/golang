package models

import (
	"api/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

var dtLocation *gorm.DB

type Location struct {
	gorm.Model
	ProvinceID     string `json:"province_id"`
	ProvinceNameth string `json:"province_nameth"`
	AmphureID      string `json:"amphure_id"`
	AmphureNameth  string `json:"amphure_nameth"`
	TambonID       string `json:"tambon_id"`
	TambonNameth   string `json:"tambon_nameth"`
	Zipcode        string `json:"zipcode"`
}

func init() {
	config.Connect()
	dtLocation = config.GetDB()
	dtLocation.AutoMigrate(&Location{})
}

func GetProvinces() []Location {
	var Provinces []Location
	dtLocation.Select("province_id, province_nameth").Group("province_id, province_nameth").Find(&Provinces)
	return Provinces
}

func GetAmphureByProvince(ID string) ([]Location, error) {
	var amphures []Location
	result := dtLocation.Select("MIN(province_id) AS province_id, amphure_id, amphure_nameth").Group("amphure_id, amphure_nameth").Where("province_id = ?", ID).Find(&amphures)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no amphures found for province ID: %s", ID)
	}
	return amphures, nil
}

func GetTambonByAmphure(ID string) ([]Location, error) {
	var tambons []Location
	result := dtLocation.Select("MIN(amphure_id) As amphure_id, tambon_id, tambon_nameth").Group("tambon_id, tambon_nameth").Where("amphure_id = ?", ID).Find(&tambons)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no tambon found for amphure ID %s", ID)
	}

	return tambons, nil
}

func GetZipcodeByTambon(ID string) ([]Location, error) {
	var zipcode []Location
	result := dtLocation.Select("tambon_id, Zipcode").Where("tambon_id = ?", ID).First(&zipcode)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no amphure found for tambon ID %s", ID)
	}

	return zipcode, nil

}
