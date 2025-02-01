package controller

import (
	"net/http"
<<<<<<< HEAD
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
=======
	"ViewUser/model"
	"github.com/gin-gonic/gin"
)

type ViewController struct{}

// Método para obtener un usuario por su ID
func (vc ViewController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := model.GetUserByID(id)
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
<<<<<<< HEAD

	// Return the user data
	c.JSON(http.StatusOK, user)
}
=======
	c.JSON(http.StatusOK, user)
}

// Método para listar todos los usuarios
func (vc ViewController) GetAllUsers(c *gin.Context) {
	users, err := model.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si no hay usuarios, devolvemos un mensaje vacío
	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No users found"})
		return
	}

	c.JSON(http.StatusOK, users)
}

>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
