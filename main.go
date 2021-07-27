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
	user.POST("/login", controllers.Login)

	user.Use(middlewares.IsAuth)
	user.GET("/", controllers.User)

	api.Use(middlewares.IsAuth)

	company := api.Group("company")
	company.GET("/",controllers.Companies)
	company.POST("/",controllers.CreateCompany)
	company.GET("/:id",controllers.GetCompany)
	company.PUT("/:id",controllers.UpdateCompany)
	company.DELETE("/:id",controllers.DeleteCompany)

	genericName := api.Group("generic")
	genericName.GET("/",controllers.GenericNames)
	genericName.POST("/",controllers.CreateGenericName)
	genericName.GET("/:id",controllers.GetGenericName)
	genericName.PUT("/:id",controllers.UpdateGenericName)
	genericName.DELETE("/:id",controllers.DeleteGenericName)

	drug := api.Group("drug")
	drug.GET("/",controllers.Drugs)
	drug.POST("/",controllers.CreateDrug)
	drug.GET("/:id",controllers.GetDrug)
	drug.PUT("/:id",controllers.UpdateDrug)
	drug.DELETE("/:id",controllers.DeleteDrug)

	customer := api.Group("customer")
	customer.GET("/",controllers.Customers)
	customer.POST("/",controllers.CreateCustomer)
	customer.GET("/:id",controllers.GetCustomer)
	customer.PUT("/:id",controllers.UpdateCustomer)
	customer.DELETE("/:id",controllers.DeleteCustomer)




	app.Run(":8080")
}
