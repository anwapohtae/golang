package controllers

import (
	"api/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBook()
	c.JSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	bookDetails, err := models.GetBookById(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, bookDetails)
}

func CreateBook(c *gin.Context) {
	var createBook models.Book
	if err := c.ShouldBindJSON(&createBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := createBook.CreateBook()
	c.JSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book := models.DeleteBook(ID)
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var updateBook models.Book
	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	bookDetails, err := models.GetBookById(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	models.GetDB().Save(&bookDetails)
	c.JSON(http.StatusOK, bookDetails)
}
