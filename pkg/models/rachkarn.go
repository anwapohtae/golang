package models

import (
	"api/pkg/config"

	"gorm.io/gorm"
)

var rachkarntb *gorm.DB

type Rachkarn struct {
	gorm.Model
	Prefixth    string `json:"prefixth"`
	Firstnameth string `json:"firstnameth"`
	Lastnameth  string `json:"lastnameth"`
	Nickname    string `json:"nickname"`
	Prefixen    string `json:"prefixen"`
	Firstnameen string `json:"firstname_en"`
	Lastnameen  string `json:"lastname_en"`
	Position    string `json:"position"`
	Birthdayme  string `json:"birthdayme"`
	Religion    string `json:"religion"`
	Cardnumber  string `json:"cardnumber"`
	Blood       string `json:"blood"`

	Numberhouse     string `json:"numberhouse"`
	Village         string `json:"village"`
	Villagename     string `json:"villagename"`
	Alley           string `json:"alley"`
	Road            string `json:"road"`
	Tambon          string `json:"tambon"`
	Amphure         string `json:"amphure"`
	Province        string `json:"province"`
	Zipcode         string `json:"zipcode"`
	Telephone       string `json:"telephone"`
	Telephonemobile string `json:"telephonemobile"`
	Emailme         string `json:"emailme"`

	Fnumberhouse     string `json:"fnumberhouse"`
	Fvillage         string `json:"fvillage"`
	Fvillagename     string `json:"fvillagename"`
	Falley           string `json:"falley"`
	Froad            string `json:"froad"`
	Ftambon          string `json:"ftambon"`
	Famphure         string `json:"famphure"`
	Fprovince        string `json:"fprovince"`
	FzipCode         string `json:"fzipcode"`
	Ftelephone       string `json:"ftelephone"`
	FtelephoneMobile string `json:"ftelephonemobile"`
	Femail           string `json:"femail"`

	Status               string `json:"status"`
	Spousename           string `json:"spousename"`
	Spouselastname       string `json:"spouselastname"`
	Spouseoldlastname    string `json:"spouseoldlastname"`
	Spouseage            string `json:"spouseage"`
	Spousereligion       string `json:"spousereligion"`
	Spouseoccupation     string `json:"spouseoccupation"`
	Spouseoccupationwork string `json:"spouseoccupationwork"`
	Spouseworkprovince   string `json:"spouseworkprovince"`
	Spousetelephone      string `json:"spousetelephone"`

	Fathername            string `json:"fathername"`
	Fatherlastname        string `json:"fatherlastname"`
	Fatherbirthday        string `json:"fatherbirthday"`
	Fatherreligion        string `json:"fatherreligion"`
	Fatherage             string `json:"fatherage"`
	Fathernumberhouse     string `json:"fathernumberhouse"`
	Fathervillage         string `json:"fathervillage"`
	FathervillageName     string `json:"fathervillagename"`
	Fatheralley           string `json:"fatheralley"`
	Fatherroad            string `json:"fatherroad"`
	Fathertambon          string `json:"fathertambon"`
	Fatheramphure         string `json:"fatheramphure"`
	Fatherprovince        string `json:"fatherprovince"`
	FatherzipCode         string `json:"fatherzipcode"`
	Fathertelephone       string `json:"fathertelephone"`
	Fathertelephonemobile string `json:"fathertelephonemobile"`

	Mothername            string `json:"mothername"`
	Motherlastname        string `json:"motherlastname"`
	Motherbirthday        string `json:"motherbirthday"`
	Motherreligion        string `json:"motherreligion"`
	Motherage             string `json:"motherage"`
	Mothernumberhouse     string `json:"mothernumberhouse"`
	Mothervillage         string `json:"mothervillage"`
	MothervillageName     string `json:"mothervillagename"`
	Motheralley           string `json:"motheralley"`
	Motherroad            string `json:"motherroad"`
	Mothertambon          string `json:"mothertambon"`
	Motheramphure         string `json:"motheramphure"`
	Motherprovince        string `json:"motherprovince"`
	MotherzipCode         string `json:"motherzipcode"`
	Mothertelephone       string `json:"mothertelephone"`
	MothertelephoneMobile string `json:"mothertelephonemobile"`

	Datelaisen  string `json:"datelaisen"`
	Monthlaisen string `json:"monthlaisen"`
	Yearlaisen  string `json:"yearlaisen"`

	Userid string `json:"userid"`

	Laisen string `json:"laisen"`
}

func init() {
	config.Connect()
	rachkarntb = config.GetDB()
	rachkarntb.AutoMigrate(&Rachkarn{})
}

func (rachkarn *Rachkarn) CreateRachkarn() *Rachkarn {
	result := rachkarntb.Create((&rachkarn))
	if result.Error != nil {
		panic(result.Error)
	}
	return rachkarn
}

func GetAllRachkarn() []Rachkarn {
	var Rachkarns []Rachkarn
	rachkarntb.Find(&Rachkarns)
	return Rachkarns
}

func GetRachkarnByUserId(ID string) (*Rachkarn, error) {
	var getRachkarn Rachkarn
	result := rachkarntb.Where("userid=?", ID).First(&getRachkarn)
	if result.Error != nil {
		return nil, result.Error
	}
	return &getRachkarn, nil
}
