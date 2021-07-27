package controllers

import (
	"net/http"
	"project/pharmacy_api/src/database"
	"project/pharmacy_api/src/models"
	"project/pharmacy_api/src/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var data map[string]string

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
		return
	}

	

	if data["password"] != data["password_confirmed"] {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "password not match",
		})
		return
	}

	if len(data["password"]) <6{
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "password too short",
		})
		return
	} 

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	user.SetPassword(data["password"])

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
		return
	}

	var user models.User

	result := database.DB.Where("email = ?", data.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no record found",
		})
		return
	}

	if err := user.ComparePassword(data.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email of password",
		})
		return
	}

	token, err := user.JwtToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid",
		})
		return
	}

	host := utils.GetEnv("HOST", "localhost")

	c.SetCookie("jwt", token, 24*3600, "/", host, false, false)

	c.JSON(http.StatusOK, user)
}

func User(c *gin.Context) {

	Id,err := utils.GetUserId(c)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":err.Error(),
		})
		return
	}

	var user models.User

	database.DB.Where("id = ?",Id).First(&user)

	c.JSON(http.StatusOK, user)
}
