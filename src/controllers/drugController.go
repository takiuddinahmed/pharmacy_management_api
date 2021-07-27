package controllers

import (
	"net/http"
	"project/pharmacy_api/src/database"
	"project/pharmacy_api/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Drugs(c *gin.Context) {
	var drugs []models.Drug
	// var drug models.Drug
	// database.DB.Model(&drug).Find(&drugs)
	database.DB.Preload("Company").Preload("GenericName").Find(&drugs)

	c.JSON(http.StatusOK, drugs)
}

func CreateDrug(c *gin.Context) {
	var drug models.Drug

	if err := c.ShouldBindJSON(&drug); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := database.DB.Create(&drug)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, drug)

}

func GetDrug(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	drug := models.Drug{
		Id: uint(id),
	}

	result := database.DB.Preload("Company").Preload("GenericName").First(&drug)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, 
			gin.H{
				"message":result.Error,})
		return
	}

	c.JSON(http.StatusOK, drug)
}

func UpdateDrug(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	drug := models.Drug{
		Id: uint(id),
	}
	c.ShouldBindJSON(&drug)
	result := database.DB.Model(&drug).Updates(&drug)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, drug)
}

func DeleteDrug(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	drug := models.Drug{
		Id: uint(id),
	}

	result := database.DB.Delete(&drug)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
	}

}
