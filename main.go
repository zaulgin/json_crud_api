package main

import (
	"github/zaulgin/json_crud_api/controllers"
	"github/zaulgin/json_crud_api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
