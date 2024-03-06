package models

import (
	"api/pkg/config"

	"gorm.io/gorm"
)

var stafftb *gorm.DB

type Staff struct {
	Prefixid      string `json:"prefixid"`
	Prefix_nameth string `json:"prefix_nameth"`
	Nameth        string `json:"nameth"`
	Nickname      string `json:"nickname"`
	Prefix_nameen string `json:"prefix_nameen"`
	Nameen        string `json:"nameen"`
	Positionid    string `json:"positionid"`
	Positionname  string `json:"positionname"`
	Birthday      string `json:"birthday"`
	Religion      string `json:"religion"`
	Cid           string `json:"cid"`
	Bloodgroup    string `json:"bloodgroup"`

	Housenumber     string `json:"housenumber"`
	Village         string `json:"village"`
	Villagename     string `json:"villagename"`
	Alley           string `json:"alley"`
	Road            string `json:"road"`
	Tambon          string `json:"tambon"`
	Amphure         string `json:"amphure"`
	Province        string `json:"province"`
	Zipcode         string `json:"zipcode"`
	Telephone       string `json:"telephone"`
	Mobiletelephone string `json:"mobiletelephone"`
	Email           string `json:"email"`

	Fhousenumber     string `json:"fhousenumber"`
	Fvillage         string `json:"fvillage"`
	Fvillagename     string `json:"fvillagename"`
	Falley           string `json:"falley"`
	Froad            string `json:"froad"`
	Ftambon          string `json:"ftambon"`
	Famphure         string `json:"famphure"`
	Fprovince        string `json:"fprovince"`
	Fzipcode         string `json:"fzipcode"`
	Ftelephone       string `json:"ftelephone"`
	Fmobiletelephone string `json:"fmobiletelephone"`
	Femail           string `json:"femail"`

	Spousename       string `json:"spousename"`
	Spouselastname   string `json:"spouselastname"`
	Agespouse        string `json:"agespouse"`
	Spousereligion   string `json:"spousereligion"`
	Spouseoccupation string `json:"spouseoccupation"`
	Placeofwork      string `json:"placeofwork"`
	Spouseprovince   string `json:"spouseprovince"`
	Spousetelephone  string `json:"spousetelephone"`

	Fathername     string `json:"fathername"`
	Fatherlastname string `json:"fatherlastname"`
	Birthdayfather string `json:"birthdayfather"`
	Fatherreligion string `json:"fatherreligion"`
	Fatherage      string `json:"fatherage"`

	Mothername     string `json:"mothername"`
	Motherlastname string `json:"motherlastname"`
	Birthdaymother string `json:"birthdaymother"`
	Motherreligion string `json:"motherreligion"`
	Motherage      string `json:"motherage"`

	Emername            string `json:"emername"`
	Emergencylastname   string `json:"emergencylastname"`
	Relation            string `json:"relation"`
	Emerhousenumber     string `json:"emerhousenumber"`
	Emervillage         string `json:"emervillage"`
	Emervillagename     string `json:"emervillagename"`
	Emeralley           string `json:"emeralley"`
	Emerroad            string `json:"emerroad"`
	Emertambon          string `json:"emertambon"`
	Emeramphure         string `json:"emeramphure"`
	Emerprovince        string `json:"emerprovince"`
	Emerzipcode         string `json:"emerzipcode"`
	Emertelephone       string `json:"emertelephone"`
	Emertelephonemobile string `json:"emertelephonemobile"`
	Emeremail           string `json:"emeremail"`
}


func init() {
	config.Connect()
	stafftb = config.GetDB()
	stafftb.AutoMigrate(&Staff{})
}