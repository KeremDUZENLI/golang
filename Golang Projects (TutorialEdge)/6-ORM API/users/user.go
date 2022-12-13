package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

// Postgres Info
// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "test"
// )

var database *gorm.DB

type User struct {
	gorm.Model
	Name  string
	Email string
}

func CreateDatabase() {
	fmt.Println("CreateDatabase")

	file, _ := os.Create("database_sqlite3.db")
	file.Close()
}

func SendDatabase() {
	fmt.Println("SendDatabase")

	// postgresInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// database, _ = gorm.Open("postgres", postgresInfo)

	database, _ = gorm.Open("sqlite3", "database_sqlite3.db")
	defer database.Close()

	database.AutoMigrate(&User{})
}

func AllUser(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "All Users")

	// postgresInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// database, _ = gorm.Open("postgres", postgresInfo)

	database, _ := gorm.Open("sqlite3", "database_sqlite3.db")
	defer database.Close()

	var User_Array []User
	database.Find(&User_Array)
	json.NewEncoder(writer).Encode(User_Array)
}

func NewUser(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "User Created")

	database, _ := gorm.Open("sqlite3", "database_sqlite3.db")
	defer database.Close()

	url_variables := mux.Vars(reader)
	name := url_variables["name"]
	email := url_variables["email"]

	database.Create(&User{Name: name, Email: email})
}

func DeleteUser(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "User Deleted")

	database, _ := gorm.Open("sqlite3", "database_sqlite3.db")
	defer database.Close()

	url_variables := mux.Vars(reader)
	name := url_variables["name"]

	var User_Struct User
	database.Where("name = ?", name).Find(&User_Struct)
	database.Delete(&User_Struct)
}

func UpdateUser(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "User Updated")

	database, _ := gorm.Open("sqlite3", "database_sqlite3.db")
	defer database.Close()

	url_variables := mux.Vars(reader)
	name := url_variables["name"]
	email := url_variables["email"]

	var User_Struct User
	database.Where("name = ?", name).Find(&User_Struct)

	User_Struct.Email = email

	database.Save(&User_Struct)
}
