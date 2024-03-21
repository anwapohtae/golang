package controllers

import (
	"api/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRachkarn(c *gin.Context) {
	var createRachkarn models.Rachkarn
	if err := c.ShouldBindJSON(&createRachkarn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rachkarn := createRachkarn.CreateRachkarn()
	c.JSON(http.StatusOK, rachkarn)
}

func GetRachkarn(c *gin.Context) {
	rachkarns := models.GetAllRachkarn()
	c.JSON(http.StatusOK, rachkarns)
}

func GetRachkarnByUserId(c *gin.Context) {
    userId := c.Param("userid")
    ID, err := strconv.Atoi(userId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID"})
        return
    }

    rachkarnDetails, err := models.GetRachkarnByUserId(strconv.Itoa(ID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Document rachkarn Not found"})
        return
    }

    c.JSON(http.StatusOK, rachkarnDetails)
}

