package database

import (
	"fmt"
	"log"

	"postgre-project/common/env"
	"postgre-project/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

func ConnectDB() *gorm.DB {
	if instance != nil {
		return instance
	}

	url := dbUrl()
	db := gormOpen(url)

	instance = db
	return instance
}

func CloseDB() {
	db, err := instance.DB()
	if err != nil {
		log.Fatal("can not get instance")
	}

	if err = db.Close(); err != nil {
		log.Fatal("can not close database")
	}
}

func LoadDB() {
	ConnectDB().AutoMigrate(&model.Tables{})
}

func dbUrl() string {
	url_db := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_DBNAME, env.DB_PORT, env.DB_SSL)

	return url_db
}

func gormOpen(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("can not connect to database")
	}

	return db
}
