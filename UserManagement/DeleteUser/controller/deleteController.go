package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"DeleteUser/model"
)

type DeleteController struct {}

func (dc DeleteController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := model.DeleteUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}