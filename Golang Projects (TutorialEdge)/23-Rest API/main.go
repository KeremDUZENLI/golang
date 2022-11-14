package main

import (
	"fiber/book"
	"fiber/database"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Database() {
	database.DatabaseConnection, _ = gorm.Open("sqlite3", "books.db")
	fmt.Println("Database Connected!")

	database.DatabaseConnection.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated!")
}

func Routes(i *fiber.App) {
	i.Get("/book/", book.BookGet)
	i.Get("/book/:id", book.BookGet1)
	i.Post("/book/", book.BookNew)
	i.Delete("/book/:id", book.BookDelete)
}

func main() {
	fiber := fiber.New()

	Database()
	defer database.DatabaseConnection.Close()

	Routes(fiber)
	fiber.Listen(3000)
}
