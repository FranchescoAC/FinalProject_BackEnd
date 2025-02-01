package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/dgrijalva/jwt-go"
	"DeleteUser/model"
	"time"
)

var secretKey = []byte("mySecretKey")

type DeleteController struct{}

// Method to delete a user by ID
func (dc DeleteController) DeleteUser(c *gin.Context) {
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

	// Get the ID of the user to delete from the URL
	id := c.Param("id")

	// Check if the user is trying to delete their own account
	if userId != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this user"})
		return
	}

	// Delete the user from the database
	err = model.DeleteUserByID(id)
=======
	"DeleteUser/model"
)

type DeleteController struct {}

func (dc DeleteController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := model.DeleteUserByID(id)
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
<<<<<<< HEAD

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
=======
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
