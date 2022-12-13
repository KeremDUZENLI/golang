package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, _ := db.Query("SELECT * FROM album WHERE artist = ?", name)
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album

		rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		albums = append(albums, alb)
	}

	return albums, nil
}

func albumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)

	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	result, _ := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)

	id, _ := result.LastInsertId()

	return id, nil
}

func main() {
	// .ENV
	godotenv.Load("database-access.env")

	// Connection
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Addr:   os.Getenv("MYSQL_PORTS"),
		Net:    os.Getenv("MYSQL_NET"),
	}

	// Database Connected
	db, _ = sql.Open("mysql", cfg.FormatDSN())
	fmt.Println("Connected!")

	// Database Scanned by Name
	albums, _ := albumsByArtist("John Coltrane")
	fmt.Printf("Albums found: %v\n", albums)

	// Database Scanned by ID
	alb, _ := albumByID(2)
	fmt.Printf("Album found: %v\n", alb)

	// Database Add Row
	albID, _ := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	fmt.Printf("ID of added album: %v\n", albID)
}
