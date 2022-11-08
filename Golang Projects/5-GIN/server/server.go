package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price int8   `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Title_1", Price: 10},
	{ID: "2", Title: "Title_2", Price: 20},
	{ID: "3", Title: "Title_3", Price: 30},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		fmt.Println(a.ID)
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
