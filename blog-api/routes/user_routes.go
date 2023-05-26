package routes

import (
	"blog-api/controller"
	"blog-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeUserRoutes(router *gin.Engine, db *gorm.DB) {

	// Set the database instance in Gin context
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// GET handler to list all the Users
	router.GET("/api/v1/users", controller.GetAllUser)

	// GET handler to find the user by ID
	router.GET("/api/v1/users/:id", controller.GetUserByID)

	// POST handler to signup as a new user
	router.POST("/api/v1/signup", controller.SignUp)

	// POST handler to login
	router.POST("/api/v1/login", controller.Login)

	// GET handler to login
	router.GET("/api/v1/logout", middleware.RequiredAuth, controller.Logout)

	// PUT handler to update a user
	router.PUT("/api/v1/users/:id", middleware.RequiredAuth, controller.UpdateUser)

	// DELETE handler to delete a user
	router.DELETE("/api/v1/users/:id", controller.DeleteUser)

}
