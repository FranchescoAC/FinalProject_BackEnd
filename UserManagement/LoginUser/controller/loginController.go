package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"LoginUser/model"
)

type LoginController struct {}

func (lc LoginController) LoginUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Authenticate user
	if isValid := user.Authenticate(); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}