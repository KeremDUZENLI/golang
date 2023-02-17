package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() {
	_, err := gorm.Open(mysql.Open("http:http@tcp(127.0.0.1:8002)/MYSQL_DATABASE"), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}
}
