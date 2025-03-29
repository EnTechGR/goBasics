package main

import (
    "net/http"
    "github.com/gin-gonic/gin"

)

type album struct {
    ID    string   `json:"id"`
    Title string   `json:"title"`
    Artist string  `json:"artist"`
    Price  string  `json:"price"`
}

var albums = []album {
    {ID: "1", Title: "Album 1", Artist: "Artist 1", Price: "10.00"},
    {ID: "2", Title: "Album 2", Artist: "Artist 2", Price: "15.00"},
    {ID: "3", Title: "Album 3", Artist: "Artist 3", Price: "20.00"},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbum(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbum)
    router.POST("/albums", postAlbum)
    router.Run("localhost:8080")
}