package variable

import (
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

type User struct {
	gorm.Model
	Name string `json:"name"`
}

var User_Test User

type User_Array []User

var User_Array_Test []User
