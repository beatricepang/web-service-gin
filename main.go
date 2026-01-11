package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// INITIALISING THE STRUCTURE
type album struct {
	ID     string  `json: "id"`
	Title  string  `json:"title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

// album slice to seed record to album data

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vuaghan", Price: 39.99},
}

// creates JSON from the slice of album structs
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

// go doesnt enforce the order in which you dexlare functions
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind the received JSON to
	// new album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}
