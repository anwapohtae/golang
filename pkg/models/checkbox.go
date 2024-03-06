package models

import (
	"api/pkg/config"

	"gorm.io/gorm"
)

var checkboxtb *gorm.DB

type Checkbox struct {
	gorm.Model
	Type_id   string `json:"type_id"`
	Type_name string `json:"type_name"`
	Name_th   string `json:"name_th"`
	Name_en   string `json:"name_en"`
}

func init() {
	config.Connect()
	checkboxtb = config.GetDB()
	checkboxtb.AutoMigrate(&Checkbox{})
}

func (checkbox *Checkbox) CreateCheckbox() *Checkbox {
	result := checkboxtb.Create(&checkbox)
	if result.Error != nil {
		panic(result.Error)
	}
	return checkbox
}

func GetGender() []Checkbox {
	var getGenders []Checkbox
	checkboxtb.Select("id, type_id, name_th, name_en").Where("type_name = ?", "gender").Find(&getGenders)
	return getGenders
}

func GetPrefix() []Checkbox {
	var getPrefixs []Checkbox
	checkboxtb.Select("id, type_id, name_th, name_en").Where("type_name = ?", "prefix").Find(&getPrefixs)
	return getPrefixs
}

func GetPosition()  []Checkbox {
	var GetPosition []Checkbox
	checkboxtb.Select("id, type_id, name_th").Where("type_name = ?", "position").Find(&GetPosition)
	return GetPosition
}

func GetStatus() []Checkbox {
	var GetStatus []Checkbox
	checkboxtb.Select("id, type_id, name_th").Where("type_name = ?", "status").Find(&GetStatus)
	return GetStatus
}
