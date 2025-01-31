package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"registerUser/model"
	"os"
)

// Genera un token JWT
func GenerateJWT(user model.User) (string, error) {
	// Define los claims (datos del token)
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // El token expirará en 72 horas
	}

	// Crea el token usando el algoritmo HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firma el token con una clave secreta
	secretKey := os.Getenv("JWT_SECRET_KEY") // Se puede definir en variables de entorno
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}