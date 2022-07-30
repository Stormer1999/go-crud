package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stormer1999/go-crud/initializers"
	"github.com/stormer1999/go-crud/models"
)

var body struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func PostsCreate(c *gin.Context) {
	// get data

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	initializers.DB.First(&post, id)

	c.Bind(&body)

	initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	initializers.DB.First(&post, id)
	initializers.DB.Delete(&post)
	c.JSON(200, gin.H{
		"post": post,
	})
}
