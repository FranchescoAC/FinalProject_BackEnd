package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/dgrijalva/jwt-go"
	"LoginUser/model"
	"time"
)

var secretKey = []byte("mySecretKey")

type LoginController struct{}

// Método para hacer login del usuario
func (lc LoginController) LoginUser(c *gin.Context) {
	var user model.User
	// Parseamos el JSON enviado en la solicitud al modelo User
=======
	"LoginUser/model"
)

type LoginController struct {}

func (lc LoginController) LoginUser(c *gin.Context) {
	var user model.User
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
<<<<<<< HEAD

	// Verificamos si las credenciales son correctas usando el método Authenticate del modelo
=======
	// Authenticate user
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
	if isValid := user.Authenticate(); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
<<<<<<< HEAD

	// Si las credenciales son correctas, generamos el token JWT
	// Utilizamos el correo como el identificador en el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userEmail": user.Email,  // Usamos el correo como identificador
		"exp":      time.Now().Add(time.Hour * 1).Unix(),  // Token con duración de 1 hora
	})

	// Firmamos el token con nuestra clave secreta
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Retornamos el token al cliente
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}
=======
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
>>>>>>> e6b002ebf0161aa606a31e93b52720ae9ae4adad
