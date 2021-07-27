package database

import (
	"fmt"
	"project/pharmacy_api/src/models"
	"project/pharmacy_api/src/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var err error

	user := utils.GetEnv("DB_USERNAME", "")
	pass := utils.GetEnv("DB_PASSWORD", "")
	host := utils.GetEnv("DB_HOST", "")
	port := utils.GetEnv("DB_PORT", "")
	name := utils.GetEnv("DB_NAME", "")

	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,pass,host,port,name)
	fmt.Println(connectionURL)
	DB, err = gorm.Open(mysql.Open(connectionURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func AutoMigrate() {
	DB.AutoMigrate(
		models.User{},
		models.Company{},
		models.GenericName{},
		models.Drug{},
		models.Customer{},
	)
}
