package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stormer1999/go-crud/initializers"
	"github.com/stormer1999/go-crud/models"
)

var body struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func AddNewPost(c *gin.Context) {
	// get data from body
	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}

	// return success response with created data
	c.JSON(200, post)
}

func GetPosts(c *gin.Context) {
	posts := []models.Post{}
	initializers.DB.Find(&posts)

	// return query data
	c.JSON(200, posts)
}

func GetPostById(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	post := models.Post{}
	result := initializers.DB.First(&post, id)

	// id not found
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}

	// return query data
	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	post := models.Post{}
	initializers.DB.First(&post, id)

	c.Bind(&body)

	initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		})

	// return updated data
	c.JSON(200, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	initializers.DB.First(&post, id)
	initializers.DB.Delete(&post)

	// return deleted data
	c.JSON(200, post)
}
