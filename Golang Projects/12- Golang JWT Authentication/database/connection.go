package database

import (
	"jwtauth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("admin:admin@/yt_go_auth"), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect database!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
