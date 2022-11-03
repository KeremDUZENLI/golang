package database

import (
	"os"
	"server/variable"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// // MySQL
// func DATABASE() {
// 	// Database Connection
// 	database, _ := sql.Open("mysql", "database_username:database_password@tcp(127.0.0.1:9998)/MYSQL_DATABASE")
// 	defer database.Close()

// 	// Database Create Table
// 	create, _ := database.Query("CREATE TABLE if not exists database_table (name VARCHAR(100) )")
// 	defer create.Close()

// 	// Database Insert Table
// 	insert, _ := database.Query("INSERT INTO database_table VALUES('AAA')")
// 	defer insert.Close()

// 	// Database Return JSON
// 	result, _ := database.Query("Select * FROM database_table")
// 	defer result.Close()
// }

var database *gorm.DB

func CreateDatabase() {
	file, _ := os.Create("HTTP_Simple_Server.db")
	file.Close()
}

func SendDatabase() {
	database, _ = gorm.Open("sqlite3", "HTTP_Simple_Server.db")
	defer database.Close()

	database.AutoMigrate(&variable.User{})
}
