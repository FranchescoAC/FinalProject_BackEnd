package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"UpdateUser/model"  // Solo importa el paquete que usas, en este caso model
)

var secretKey = []byte("mySecretKey")

type UpdateController struct{}

// Método para actualizar un usuario por ID
func (uc UpdateController) UpdateUser(c *gin.Context) {
	// Obtener JWT del encabezado de la solicitud
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	// Validar el JWT y manejar el error correctamente
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Asegurarse de que el método de firma sea el esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// Si el token no es válido o hay un error
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Obtener el ID del usuario desde los claims del token
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)

	// Obtener el ID del usuario a actualizar desde el parámetro de la URL
	id := c.Param("id")

	// Verificar que el usuario esté intentando actualizar su propia cuenta
	if userId != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this user"})
		return
	}

	// Obtener los datos del usuario a actualizar desde el cuerpo de la solicitud
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamar a la función que actualiza el usuario en la base de datos
	updatedUser, err := model.UpdateUserByID(userId, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver el usuario actualizado
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": updatedUser})
}
