package controller

import (
	"blog-api/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Create a user to bind with the request body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Generate a hash from the given password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash the password"})
		return
	}
	user.Password = string(hash)

	// Create the user in the database
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Return the created user
	c.JSON(200, user)

}

func Login(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Create a user to bind with the request body
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get the actual user from the database
	var actualUser models.User
	db.First(&actualUser, "email=?", user.Email)
	if actualUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or Password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(actualUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or Password"})
		return
	}

	//gebrate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": actualUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// SignIn and get the token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	// Set the coockie with the generated token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Cookiee is set.",
	})

}

func Logout(c *gin.Context) {
	// Clear the JWT token cookie
	c.SetCookie("token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func GetAllUser(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Declare an array of user to save the result
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, users)
}

func GetUserByID(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get the user ID from the URL parameter
	userID := c.Param("id")

	var user models.User
	result := db.Find(&user, userID)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	if user.ID == 0 {
		c.JSON(404, gin.H{"Error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get the user ID from the URL parameter
	userID := c.Param("id")

	// Bind JSON request body to User struct
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update the user in the database
	result := db.Model(&models.User{}).Where("id = ?", userID).Updates(updatedUser)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	} else if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"Error": "User not found"})
		return
	} else {
		var user models.User
		result := db.Find(&user, userID)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, user)
		return
	}
}

func DeleteUser(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get the user ID from the URL parameter
	userID := c.Param("id")

	// Delete the user from the database
	result := db.Delete(&models.User{}, userID)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"Error": "User not found"})
		return
	}

	// Return a success message
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
