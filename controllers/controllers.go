package controllers

import (
	"gin_course/initialize"
	"gin_course/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	initialize.DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums) // Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil { // Use Context.BindJSON to bind the request body to newAlbum.
		c.IndentedJSON(http.StatusBadRequest, c.Error(err))
		return
	}

	initialize.DB.Create(&newAlbum)              // Append the album struct initialized from the JSON to the albums db.
	c.IndentedJSON(http.StatusCreated, newAlbum) // Add a 201 status code to the response, along with JSON representing the album you added.
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id") // When you map this handler to a path, you will include a placeholder for the parameter in the path.
	var album models.Album

	initialize.DB.First(&album, id)

	if album.ID != "" {
		c.IndentedJSON(http.StatusOK, album)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var album models.Album

	initialize.DB.First(&album, id)

	if album.ID == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Couldn't find an album by current id"})
		return
	}

	initialize.DB.Delete(&models.Album{}, id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

func ChangeByID(c *gin.Context) {
	id := c.Param("id")

	var newalbum, album models.Album

	initialize.DB.First(&album, id)

	newalbum = album

	if album.ID == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	if err := c.Bind(&newalbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldn't unmarshal your request"})
	}

	initialize.DB.Model(&album).Updates(newalbum)
}
