package controllers

import (
	"api/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	newUsers := models.GetAllUser()
	c.JSON(http.StatusOK, newUsers)
}

func GetUserById(c *gin.Context) {
	userId := c.Param("userId")
	ID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invakid user ID"})
		return
	}

	userDetails, err := models.GetUserById(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, userDetails)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := models.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}



func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	ID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user := models.DeleteUser(ID)
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var updateUser models.User

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.Param("userId")
	ID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userDetails, err := models.GetUserById(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if updateUser.Firstname != "" {
		userDetails.Firstname = updateUser.Firstname
	}

	if updateUser.Lastname != "" {
		userDetails.Lastname = updateUser.Lastname
	}

	if updateUser.Age != "" {
		userDetails.Age = updateUser.Age
	}

	if updateUser.Numberphone != "" {
		userDetails.Numberphone = updateUser.Numberphone
	}

	if updateUser.Role != "" {
		userDetails.Role = updateUser.Role
	}

	if updateUser.Googlefirstsname != "" {
		userDetails.Googlefirstsname = updateUser.Googlefirstsname
	}

	if updateUser.Googlelastname != "" {
		userDetails.Googlelastname = updateUser.Googlelastname
	}

	if updateUser.Googleid != "" {
		userDetails.Googleid = updateUser.Googleid
	}

	if updateUser.Profile != "" {
		userDetails.Profile = updateUser.Profile
	}

	models.GetDB().Save(&userDetails)
	c.JSON(http.StatusOK, userDetails)

}
