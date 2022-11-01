package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "testes:testes@tcp(localhost:3306)/testes?parseTime=true")
	if err != nil {
		panic(err)
	}

	fmt.Printf("It is running don't worry")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
