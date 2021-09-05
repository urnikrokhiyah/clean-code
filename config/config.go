package config

import (
	"cleancode/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	// config := map[string]string{
	// 	"Db_Username": "root",
	// 	"Db_Password": "alta123",
	// 	"Db_Port":     "3306",
	// 	"Db_Host":     "127.0.0.1",
	// 	"Db_Name":     "cleancode",
	// }

	connection := os.Getenv("CONNECTION")

	var err error
	Db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {
	Db.AutoMigrate(&models.User{})
}
