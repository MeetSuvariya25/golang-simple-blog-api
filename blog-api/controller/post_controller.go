package controller

import (
	"blog-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPost(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	page := c.Query("page")
	limit := c.Query("limit")

	// Convert page and limit parameters to integers
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Set default values if page or limit are not provided
	if pageInt == 0 {
		pageInt = 1
	}
	if limitInt == 0 {
		limitInt = 10
	}

	// Calculate the offset
	offset := (pageInt - 1) * limitInt

	var posts []models.Post
	result := db.Limit(limitInt).Offset(offset).Preload("User").Find(&posts)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, posts)
}

func GetAllPostByAuthor(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get the user ID from the URL parameter
	ByAuthor := c.Param("user")

	var posts []models.Post
	result := db.Find(&posts, ByAuthor)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	if len(posts) == 0 {
		c.JSON(404, gin.H{"Error": "Posts not found for the givem author"})
		return
	}

	c.JSON(200, posts)
}

func GetPostByID(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get the user ID from the URL parameter
	postID := c.Param("id")

	var post models.Post
	result := db.Find(&post, postID)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	if post.ID == 0 {
		c.JSON(404, gin.H{"Error": "Post not found"})
		return
	}

	c.JSON(200, post)
}

func AddPost(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Bind JSON request body to User struct
	var post = &models.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create the user in the database
	result := db.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	// Return the created user
	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get the user ID from the URL parameter
	postID := c.Param("id")

	// Bind JSON request body to User struct
	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update the user in the database
	result := db.Model(&models.Post{}).Where("id = ?", postID).Updates(updatedPost)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Return the updated user
	c.JSON(200, updatedPost)
}

func DeletePost(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get the user ID from the URL parameter
	postID := c.Param("id")

	// Delete the user from the database
	result := db.Delete(&models.Post{}, postID)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Return a success message
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
