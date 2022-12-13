package book

import (
	"fiber/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func BookGet(c *fiber.Ctx) {
	table := database.DatabaseConnection

	var bookGet []Book
	table.Find(&bookGet)

	c.JSON(bookGet)
}

func BookGet1(c *fiber.Ctx) {
	table := database.DatabaseConnection

	var bookGet1 Book
	id := c.Params("id")
	table.Find(&bookGet1, id)

	c.JSON(bookGet1)
}

func BookNew(c *fiber.Ctx) {
	table := database.DatabaseConnection

	bookNew := new(Book)
	c.BodyParser(bookNew)
	table.Create(&bookNew)

	c.JSON(bookNew)
}

func BookDelete(c *fiber.Ctx) {
	table := database.DatabaseConnection

	var deleteBook Book
	id := c.Params("id")
	table.Find(&deleteBook, id)
	table.Delete(&deleteBook)

	c.JSON(deleteBook)
}
