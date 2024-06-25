package main

import (
	"gin_course/controllers"
	"gin_course/initialize"

	"github.com/gin-gonic/gin"
)

func init() {
	initialize.ConnectToDB()
}

func main() {
	router := gin.Default() // Default returns an Engine instance with the Logger and Recovery middleware already attached

	router.GET("/albums", controllers.GetAlbums)                     // Implements the method GET for the provided enpoint
	router.GET("/albums/:id", controllers.GetAlbumByID)              //In Gin, the colon preceding an item in the path signifies that the item is a path parameter.
	router.POST("/albums", controllers.PostAlbums)                   // You can separately route requests sent to a single path based on the method the client is using.
	router.DELETE("/albums/delete/:id", controllers.DeleteAlbumByID) // We delete album by id
	router.PUT("/albums/change/:id", controllers.ChangeByID)

	router.Run("localhost:8080") //Run attaches the router to a http.Server and starts listening and serving HTTP requests.
}
