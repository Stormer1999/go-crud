package main

import (
	"github.com/stormer1999/go-crud/initializers"
	"github.com/stormer1999/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
