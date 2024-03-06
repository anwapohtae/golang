package models

import (
	"api/pkg/config"

	"gorm.io/gorm"
)

var governmenttb *gorm.DB

type Government struct {
	gorm.Model
	
}

func init() {
	config.Connect()
	governmenttb = config.GetDB()
	governmenttb.AutoMigrate(&Government{})
}


