package middleware

import (
	"blog-api/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func RequiredAuth(c *gin.Context) {
	// Get the database instance from Gin context
	db := c.MustGet("db").(*gorm.DB)

	//Get the cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//Decode
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, Ok := token.Method.(*jwt.SigningMethodHMAC); !Ok {
			return nil, fmt.Errorf("Unexpected signin method : %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECERET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find user with token
		var user models.User
		db.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Attch to req
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
