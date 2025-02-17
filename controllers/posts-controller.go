package controllers

import (
	"net/http"

	"github.com/azacdev/go-crud/initializers"
	"github.com/azacdev/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create a post
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

func GetPosts(c *gin.Context) {
	// Create a posts variable (slice of Post)
	var posts []models.Post

	// Get all posts
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts, // Return the slice of posts
	})
}

func GetPost(c *gin.Context) {
	// Get id of url

	id := c.Param("id")

	// Create a posts variable (slice of Post)
	var posts models.Post

	// Get all posts
	result := initializers.DB.First(&posts, id)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts, // Return the slice of posts
	})
}

func UpdatePost(c *gin.Context) {
	// Get id of url
	id := c.Param("id")

	// get the posts
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get all posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

func DeletePost(c *gin.Context) {
	// Get id of url

	id := c.Param("id")

	// Get all posts
	result := initializers.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(200)
}
