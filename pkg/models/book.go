package models

import (
	"api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	if result.Error != nil {
		panic(result.Error)
	}
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, error) {
	var getBook Book
	result := db.Where("ID=?", ID).First(&getBook)
	if result.Error != nil {
		return nil, result.Error
	}
	return &getBook, nil
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}

func GetDB() *gorm.DB {
	return db
}
