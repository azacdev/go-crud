package main

import (
	"github.com/azacdev/go-crud/controllers"
	"github.com/azacdev/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	initializers.ConnectToDB()
	r := gin.Default()

	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
