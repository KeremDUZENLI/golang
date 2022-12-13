package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func DatabaseBegin() {
	database_http, _ := gorm.Open("mysql", "http:http@tcp(127.0.0.1:8001)/MYSQL_DATABASE")

	fmt.Printf("Database Begin")
	database = database_http
}

func DatabaseGet() *gorm.DB {
	return database
}
