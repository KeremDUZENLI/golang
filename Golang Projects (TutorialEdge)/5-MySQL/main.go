package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	// Create database
	database, fail := sql.Open("mysql", "kerem:kerem@tcp(127.0.0.1:3306)/MYSQL_1")

	if fail != nil {
		panic(fail.Error())
	}

	defer database.Close()

	// Insert database
	insert, fail := database.Query("INSERT INTO users VALUES('KEREM')")

	if fail != nil {
		panic(fail.Error())
	}

	defer insert.Close()

	// JSON results
	results, fail := database.Query("SELECT name FROM users")

	if fail != nil {
		panic(fail.Error())
	}

	for results.Next() {
		var user User

		fail = results.Scan(&user.Name)
		if fail != nil {
			panic(fail.Error())
		}

		fmt.Println(user.Name)
	}
}
