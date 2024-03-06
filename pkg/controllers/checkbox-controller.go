package controllers

import (
	"api/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCheckbox(c *gin.Context) {
	var createCheckbox models.Checkbox
	if err := c.ShouldBindJSON(&createCheckbox); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkbox := createCheckbox.CreateCheckbox()
	c.JSON(http.StatusOK, checkbox)
}

func GetGender(c *gin.Context) {
	Genders := models.GetGender()
	c.JSON(http.StatusOK, Genders)
}

func GetPrefix(c *gin.Context) {
	Prefixs := models.GetPrefix()
	c.JSON(http.StatusOK, Prefixs)
}

func GetPosition(c *gin.Context) {
	Position := models.GetPosition()
	c.JSON(http.StatusOK, Position)
}

func GetStatus(c *gin.Context) {
	Status := models.GetStatus()
	c.JSON(http.StatusOK, Status)
}