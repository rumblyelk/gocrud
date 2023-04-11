package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rumblyelk/gocrud/initializers"
	"github.com/rumblyelk/gocrud/models"
)

func PostsCreate(c *gin.Context) {
	//get data off request body
	var body struct {
		Title string
		Body  string
	}
	c.BindJSON(&body)

	//create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return post
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
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// here we check if the post is not found
	if post.ID == 0 {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// here we check if the post is not found
	if post.ID == 0 {
		c.Status(404)
		return
	}

	//get data off request body
	var body struct {
		Title string
		Body  string
	}
	c.BindJSON(&body)

	//update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDestroy(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))

	// here we check if the post is not found
	if post.ID == 0 {
		c.Status(404)
		return
	}

	initializers.DB.Delete(&post)

	c.Status(200)
}
