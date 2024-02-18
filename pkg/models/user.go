package models

import (
	"api/pkg/config"

	"gorm.io/gorm"
)

var usertb *gorm.DB

type User struct {
	gorm.Model
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Googlefirstsname string `json:"googlefirstsname"`
	Googlelastname  string `json:"googlelastname"`
	Age             string `json:"age"`
	Numberphone     string `json:"numberphone"`
	Googleid        string `json:"googleid"`
	Email           string `json:"email"`
	Profile         string `json:"profile"`
	Role            string `json:"role"`
}

func init() {
	config.Connect()
	usertb = config.GetDB()
	usertb.AutoMigrate(&User{})
}

func (u *User) CreatedUser() *User {
	result := usertb.Create(&u)
	if result.Error != nil {
		panic(result.Error)
	}

	return u
}

func GetAllUser() []User {
	var Users []User
	usertb.Find(&Users)
	return Users
}

func GetUserById(ID int64) (*User, error) {
	var getUser User
	result := usertb.Where("ID=?", ID).First(&getUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &getUser, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	result := usertb.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}


func DeleteUser(ID int64) User {
	var user User
	usertb.Where("ID=?", ID).Delete(&user)
	return user
}

func GetUsertb() *gorm.DB {
	return usertb
}
