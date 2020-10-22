package main

import (
	"app/controller"
	"app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, World")
	})
	r.POST("/post", controller.CreatePost)
	r.GET("/posts", controller.GetAllPosts)

	r.Run()
}
