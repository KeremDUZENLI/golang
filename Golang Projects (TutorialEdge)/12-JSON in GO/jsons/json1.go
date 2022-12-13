package jsons

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"developer"`
}

func Json1() {
	newBook := Book{Title: "abc", Author: "Kerem"}
	fmt.Printf("%+v\n", newBook)

	byteArray1, err1 := json.Marshal(newBook)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Print(string(byteArray1), "\n\n")

	newAuthor := Author{Name: "Kerem", Age: 20, Developer: true}
	fmt.Printf("%+v\n", newAuthor)

	byteArray2, err2 := json.MarshalIndent(newAuthor, "", "	")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Print(string(byteArray2), "\n\n")
}
