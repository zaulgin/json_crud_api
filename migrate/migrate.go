package main

import (
	"github/zaulgin/json_crud_api/initializers"
	"github/zaulgin/json_crud_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
