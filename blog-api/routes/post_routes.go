package routes

import (
	"blog-api/controller"
	"blog-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializePostRoutes(router *gin.Engine, db *gorm.DB) {

	// GET handler to list all the posts by Users
	router.GET("/posts/specific/:user", controller.GetAllPostByAuthor)

	// GET handler to list all the post
	router.GET("/posts", controller.GetAllPost)

	// GET handler to list the post by ID
	router.GET("/posts/:id", controller.GetPostByID)

	// POST handler to add a new post
	router.POST("/posts", middleware.RequiredAuth, controller.AddPost)

	// PUT handler to update the post
	router.PUT("/posts/:id", middleware.RequiredAuth, controller.UpdatePost)

	// DELETE handler to delete the post
	router.DELETE("/posts/:id", middleware.RequiredAuth, controller.DeletePost)

}
