package main

import (
	"example/web-service-golang/src/useCases/albums"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albumsvar = []albums.Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albumsvar)
}

func postAlbums(c *gin.Context) {
    var newAlbum albums.Album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albumsvar = append(albumsvar, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Slow search, verify in all data if ID == userID
    for _, a := range albumsvar {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
