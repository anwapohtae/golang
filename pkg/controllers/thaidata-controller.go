package controllers

import (
	"api/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetThaiData(c *gin.Context) {
	Thaidatas := models.GetAllThaiData()
	c.JSON(http.StatusOK, Thaidatas)
}