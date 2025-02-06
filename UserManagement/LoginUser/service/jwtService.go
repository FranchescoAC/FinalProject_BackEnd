package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET_KEY = []byte("mysecretkey")

// Generar un token JWT para el usuario autenticado
func GenerateJWT(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       userID,
		"username": username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(), // Token válido por 72 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_SECRET_KEY)
}
