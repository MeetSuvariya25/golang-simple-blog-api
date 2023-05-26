package routes

import (
	"blog-api/controller"
	"blog-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializePostRoutes(router *gin.Engine, db *gorm.DB) {

	// GET handler to list all the posts by Users
	router.GET("/api/v1/posts/specific/:user", controller.GetAllPostByAuthor)

	// GET handler to list all the post
	router.GET("/api/v1/posts", controller.GetAllPost)

	// GET handler to list the post by ID
	router.GET("/api/v1/posts/:id", controller.GetPostByID)

	// POST handler to add a new post
	router.POST("/api/v1/posts", middleware.RequiredAuth, controller.AddPost)

	// PUT handler to update the post
	router.PUT("/api/v1/posts/:id", middleware.RequiredAuth, controller.UpdatePost)

	// DELETE handler to delete the post
	router.DELETE("/api/v1/posts/:id", middleware.RequiredAuth, controller.DeletePost)

}
