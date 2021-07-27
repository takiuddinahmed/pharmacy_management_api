package controllers

import (
	"net/http"
	"project/pharmacy_api/src/database"
	"project/pharmacy_api/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Customers(c *gin.Context) {
	var customers []models.Customer

	database.DB.Find(&customers)

	c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := database.DB.Create(&customer)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, customer)

}

func GetCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer := models.Customer{
		Id: uint(id),
	}

	result := database.DB.First(&customer)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, 
			gin.H{
				"message":result.Error,})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer := models.Customer{
		Id: uint(id),
	}
	c.ShouldBindJSON(&customer)
	result := database.DB.Model(&customer).Updates(&customer)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer := models.Customer{
		Id: uint(id),
	}

	result := database.DB.Delete(&customer)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
	}

}
