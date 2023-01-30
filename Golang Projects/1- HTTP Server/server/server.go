package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var User_Test User

type User_Array []User

var User_Array_Test User_Array = []User{}

func BEGINURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "HOMEPAGE")
}

func GETURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "GET URL\n")

	json.NewEncoder(writer).Encode(User_Array_Test)
}

func POSTURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "POST URL\n")

	json.NewDecoder(reader.Body).Decode(&User_Test)
	json.NewEncoder(writer).Encode(User_Test)

	User_Array_Test = append(User_Array_Test, User_Test)

	writer.Header().Set("Content-Type", "application/json")
}

func DELETEURL(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "DELETE URL\n")

	url := reader.RequestURI[7:]
	fmt.Println(url)

	for i := 0; i < len(User_Array_Test); i++ {
		if url == User_Array_Test[i].Name {
			User_Array_Test = append(User_Array_Test[:i], User_Array_Test[i+1:]...)
		}
	}
}
