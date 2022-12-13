package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {
	// Open up db connection
	db, err := sql.Open("mysql", "testes:testes@tcp(db:3306)/testes")

	// check if error with db connection
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Execute query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var users []*User
	for results.Next() {
		var u User

		// scan each result
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}

		users = append(users, &u)
	}

	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("Hit the home page endpoint")
}

// usersPage
func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Hit the users page endpoint")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
