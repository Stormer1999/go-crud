package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stormer1999/go-crud/controllers"
	"github.com/stormer1999/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.POST("/posts", controllers.AddNewPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run() // listen and serve on
}
