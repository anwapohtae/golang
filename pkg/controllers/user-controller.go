package controllers

import (
	"api/pkg/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func CreateUser(c *gin.Context) {
	var createUser models.User
	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้าง JWT Claims
	claims := jwt.MapClaims{
		"firstname": createUser.Firstname,
		"email":     createUser.Email,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // กำหนดวันหมดอายุของ Token เป็น 24 ชั่วโมง
	}

	// สร้าง Token ด้วยความลับ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key")) // คีย์ลับที่คุณเลือกใช้

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	u := createUser.CreatedUser()
	c.JSON(http.StatusOK, gin.H{"user create successfully": u, "token": tokenString})
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

func Login(c *gin.Context) {
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user data from the database using the provided email
	var userFromDB models.User
	err := models.GetUsertb().Where("email = ?", loginUser.Email).First(&userFromDB).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "อีเมลนี้ยังไม่มีในระบบ"})
		return
	}

	// Update the user details if the email matches
	if loginUser.Email == userFromDB.Email {
		// Update user details based on the provided information
		if loginUser.Firstname != "" {
			userFromDB.Firstname = loginUser.Firstname
		}
		if loginUser.Lastname != "" {
			userFromDB.Lastname = loginUser.Lastname
		}
		if loginUser.Age != "" {
			userFromDB.Age = loginUser.Age
		}
		if loginUser.Numberphone != "" {
			userFromDB.Numberphone = loginUser.Numberphone
		}
		if loginUser.Role != "" {
			userFromDB.Role = loginUser.Role
		}
		if loginUser.Googlefirstsname != "" {
			userFromDB.Googlefirstsname = loginUser.Googlefirstsname
		}
		if loginUser.Googlelastname != "" {
			userFromDB.Googlelastname = loginUser.Googlelastname
		}
		if loginUser.Googleid != "" {
			userFromDB.Googleid = loginUser.Googleid
		}
		if loginUser.Profile != "" {
			userFromDB.Profile = loginUser.Profile
		}

		// Save the updated user details back to the database
		if err := models.GetDB().Save(&userFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user details"})
			return
		}

		// สร้าง JWT Claims
		claims := jwt.MapClaims{
			"userId":           userFromDB.ID,
			"firstname":        userFromDB.Firstname,
			"lastname":         userFromDB.Lastname,
			"googlefirstsname": userFromDB.Googlefirstsname,
			"googlelastname":   userFromDB.Googlelastname,
			"age":              userFromDB.Age,
			"numberphone":      userFromDB.Numberphone,
			"googleid":         userFromDB.Googleid,
			"email":            userFromDB.Email,
			"profile":          userFromDB.Profile,
			"role":             userFromDB.Role,
			"exp":              time.Now().Add(time.Hour * 24).Unix(), // กำหนดวันหมดอายุของ Token เป็น 24 ชั่วโมง
		}

		// สร้าง Token ด้วยความลับ
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("your-256-bit-secret")) // คีย์ลับที่คุณเลือกใช้

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}

		// ส่งข้อมูลผู้ใช้และ Token กลับไปยังผู้ใช้
		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
