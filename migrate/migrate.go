package main

import (
	"crud-gin-gorm/initializers"
	"crud-gin-gorm/models"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}