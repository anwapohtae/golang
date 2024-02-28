package controllers

import (
	"api/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func GetAllLocations(c *gin.Context) {
// 	Provinces := models.GetAllLocations()
// 	c.JSON(http.StatusOK, Provinces)
// }

func GetProvinces(c *gin.Context) {
	Provinces := models.GetProvinces()
	c.JSON(http.StatusOK, Provinces)
}

func GetAmphureByProvince(c *gin.Context) {
    provinceId := c.Param("provinceId")

    amphureDetails, err := models.GetAmphureByProvince(provinceId)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Amphure not found"})
        return
    }

    c.JSON(http.StatusOK, amphureDetails)
}

func GetTambonByAmphure(c *gin.Context) {
    amphureId := c.Param("amphureId")

    tambonDetails, err := models.GetTambonByAmphure(amphureId)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tambon not found"})
        return
    }

    c.JSON(http.StatusOK, tambonDetails)
}

func GetZipcodeByTambon(c *gin.Context) {
    tambonId := c.Param("tambonId")

    zipcodeDetails, err := models.GetZipcodeByTambon(tambonId)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Zipcode not found"})
        return
    }

    c.JSON(http.StatusOK, zipcodeDetails)
}

// func GetAmphureByProvince(c *gin.Context) {
// 	provinceId := c.Param("provinceId")
// 	ID, err := strconv.ParseInt(provinceId, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provinces ID"})
// 		return
// 	}

// 	// แปลง ID เป็น string
// 	IDString := strconv.Itoa(int(ID))

// 	amphureDetails, err := models.GetAmphureByProvince(IDString)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Amphure not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, amphureDetails)
// }

// func GetAmphures(c *gin.Context) {
// 	Amphures := models.GetAmphures()
// 	c.JSON(http.StatusOK, Amphures)
// }

// func GetTambons(c *gin.Context) {
// 	Tambons := models.GetTambons()
// 	c.JSON(http.StatusOK, Tambons)
// }

// func GetZipcodes(c *gin.Context) {
// 	Zipcodes := models.GetZipcode()
// 	c.JSON(http.StatusOK, Zipcodes)
// }

// func SelectedProvince(c *gin.Context) {
// 	Provincedatas := models.SelectedProvince()
// 	c.JSON(http.StatusOK, Provincedatas)
// }
