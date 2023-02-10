package configs

import (
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB

func DatabaseDB() {
	database_mysql, err := gorm.Open("mysql", "test:test@tcp(127.0.0.1:3307)/MYSQL_REST_API_TEST")

	if err != nil {
		panic("NOT CONNECTED!")
	}

	Database = database_mysql
}
