package controllers

import (
	"api/pkg/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)



func Register(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "อีเมลนี้ยังไม่มีในระบบ",
		})
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