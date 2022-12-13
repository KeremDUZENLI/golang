package model

import (
	config "http/configs"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

type User struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name"`
}

func init() {
	config.DatabaseBegin()
	database = config.DatabaseGet()
	database.AutoMigrate(&User{})
}

func GetUser() *[]User {
	var userAllGet []User
	database.Find(&userAllGet)
	return &userAllGet
}

func GetUserID(ID int64) (*User, *gorm.DB) {
	var userIDGet User
	findings := database.Where("ID=?", ID).Find(&userIDGet)
	return &userIDGet, findings
}

func (userStruct *User) PostUser() *User {
	database.NewRecord(userStruct)
	database.Create(&userStruct)
	return userStruct
}

func DeleteUser(ID int64) User {
	var userDelete User
	database.Where("ID=?", ID).Delete(userDelete)
	return userDelete
}
