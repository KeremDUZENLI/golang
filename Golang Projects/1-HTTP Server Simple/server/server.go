package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/variable"

	"github.com/jinzhu/gorm"
)

func BEGINURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "HOMEPAGE")
}

func GETURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "GET URL")

	db, _ := gorm.Open("sqlite3", "HTTP_Simple_Server.db")
	defer db.Close()

	db.Find(&variable.User_Array_Test)
	json.NewEncoder(writer).Encode(variable.User_Array_Test)
}

func POSTURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "POST URL")

	db, _ := gorm.Open("sqlite3", "HTTP_Simple_Server")
	defer db.Close()

	json.NewDecoder(reader.Body).Decode(&variable.User_Test)
	json.NewEncoder(writer).Encode(variable.User_Test)

	variable.User_Array_Test = append(variable.User_Array_Test, variable.User_Test)

	writer.Header().Set("Content-Type", "application/json")

	db.Create(&variable.User{Name: variable.User_Test.Name})
}

func DELETEURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "DELETE URL")

	url := reader.RequestURI[1:]

	for i := 0; i < len(variable.User_Array_Test); i++ {
		if url == variable.User_Array_Test[i].Name {
			variable.User_Array_Test = append(variable.User_Array_Test[:i], variable.User_Array_Test[i+1:]...)
		}
	}
}
