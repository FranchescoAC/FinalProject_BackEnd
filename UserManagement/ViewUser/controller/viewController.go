package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"ViewUser/model"
	"time"
)

var secretKey = []byte("mySecretKey")

type ViewController struct{}

// Method to get a user by ID
func (vc ViewController) GetUser(c *gin.Context) {
	// Get JWT from the Authorization header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	// Validate JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, gin.Error{Err: err, Type: gin.ErrorTypePublic}
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Get the user ID from the token claims
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)

	// Fetch the user from the database
	user, err := model.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Return the user data
	c.JSON(http.StatusOK, user)
}
