package main

import (
	"gin/server"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()

	ginRouter.GET("/albums", server.GetAlbums)
	ginRouter.GET("/albums/:id", server.GetAlbumsByID)
	ginRouter.POST("/albums", server.PostAlbums)

	ginRouter.Run("localhost:8080")
}
