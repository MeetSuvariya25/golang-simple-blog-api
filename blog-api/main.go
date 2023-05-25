package main

import (
	"blog-api/config"
	"blog-api/initializers"
	"blog-api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LocalEnvVariable()
}

func main() {

	// Initialize the database
	db := config.InitializeDB()
	defer config.CloseDB(db)

	// Create a new Gin router
	router := gin.Default()

	// Initialize the API routes for users
	routes.InitializeUserRoutes(router, db)

	// Initialize the API routes for posts
	routes.InitializePostRoutes(router, db)

	// Start the server
	router.Run(":8080")
}
