package main

import (
	"project/pharmacy_api/src/controllers"
	"project/pharmacy_api/src/database"
	"project/pharmacy_api/src/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load environment variables
	godotenv.Load()

	// start database
	database.Connect()
	database.AutoMigrate()

	

	app := gin.Default()

	api := app.Group("api")

	user := api.Group("user")

	user.GET("/register", controllers.Register)
	user.GET("/login", controllers.Login)

	user.Use(middlewares.IsAuth)
	user.GET("/user", controllers.User)

	app.Run(":8080")
}
