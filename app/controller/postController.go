package controller

import (
	"app/database"
	"app/models"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	posts := models.GetAllPosts(database.GetDB())

	// for _, v := range posts {
	// 	println(v.Title)
	// }

	c.JSON(200, posts)
}

func CreatePost(c *gin.Context) {
	post := &models.Post{}
	err := c.Bind(post)
	if err != nil {
		println(err)
	}

	post.CreatePost(database.GetDB())
	// err = post.CreatePost(database.GetDB())
	// if err != nil {
	// 	println(err)
	// }

	c.JSON(200, post)
}
