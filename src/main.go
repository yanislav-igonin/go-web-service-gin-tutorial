package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.Run("localhost:" + port)
}
