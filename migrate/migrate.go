package main

import (
	"gin_course/initialize"
	"gin_course/models"
)

func init() {
	initialize.ConnectToDB()
}

func main() {
	initialize.DB.AutoMigrate(&models.Album{})
}
