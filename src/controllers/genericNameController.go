package controllers

import (
	"net/http"
	"project/pharmacy_api/src/database"
	"project/pharmacy_api/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GenericNames(c *gin.Context) {
	var genericNames []models.GenericName

	database.DB.Find(&genericNames)

	c.JSON(http.StatusOK, genericNames)
}

func CreateGenericName(c *gin.Context) {
	var genericName models.GenericName

	if err := c.ShouldBindJSON(&genericName); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := database.DB.Create(&genericName)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, genericName)

}

func GetGenericName(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	genericName := models.GenericName{
		Id: uint(id),
	}

	result := database.DB.First(&genericName)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, 
			gin.H{
				"message":result.Error,})
		return
	}

	c.JSON(http.StatusOK, genericName)
}

func UpdateGenericName(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	genericName := models.GenericName{
		Id: uint(id),
	}
	c.ShouldBindJSON(&genericName)
	result := database.DB.Model(&genericName).Updates(&genericName)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, genericName)
}

func DeleteGenericName(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	genericName := models.GenericName{
		Id: uint(id),
	}

	result := database.DB.Delete(&genericName)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
	}

}
