package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/gowebexampes/sqlc-examples/postgres"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "user=gowebexamples password=gowebexamples dbname=sqlc sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db := postgres.New(conn)

	user, err := db.CreateUser(context.Background(), postgres.CreateUserParams{
		Firstname: "Kerem",
		Lastname:  "DUZENLI",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(user)
}
