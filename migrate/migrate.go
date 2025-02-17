package main

import (
	"github.com/azacdev/go-crud/initializers"
	"github.com/azacdev/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
