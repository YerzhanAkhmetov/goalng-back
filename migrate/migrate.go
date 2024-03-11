package main

import (
	"migration/initializer"
	"migration/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}
func main() {
	initializer.DB.AutoMigrate(&models.Post{})
	// 	initializer.DB.AutoMigrate(&models.User{})
}
