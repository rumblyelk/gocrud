package main

import (
	"github.com/rumblyelk/gocrud/initializers"
	"github.com/rumblyelk/gocrud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
