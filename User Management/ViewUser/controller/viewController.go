package controller

import (
	"net/http"
	"ViewUser/model"
	"github.com/gin-gonic/gin"
)

type ViewController struct{}

// Método para obtener un usuario por su ID
func (vc ViewController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := model.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
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

