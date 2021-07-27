package controllers

import (
	"net/http"
	"project/pharmacy_api/src/database"
	"project/pharmacy_api/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Companies(c *gin.Context) {
	var companies []models.Company

	database.DB.Find(&companies)

	c.JSON(http.StatusOK, companies)
}

func CreateCompany(c *gin.Context) {
	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := database.DB.Create(&company)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, company)

}

func GetCompany(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	company := models.Company{
		Id:uint(id),
	}

	database.DB.Find(&company)

	c.JSON(http.StatusOK,company)
}	

func UpdateCompany(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))

	company := models.Company{
		Id:uint(id),
	}
	c.ShouldBindJSON(&company)
	result :=database.DB.Model(&company).Updates(&company)

	if result.Error!= nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message":result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, company)
}

func DeleteCompany(c *gin.Context){
	id,_:= strconv.Atoi(c.Param("id"))

	company := models.Company{
		Id: uint(id),
	}

	result := database.DB.Delete(&company)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"message": result.Error,
		})	
	}

}
