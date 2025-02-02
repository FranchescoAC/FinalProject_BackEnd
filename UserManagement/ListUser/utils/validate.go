package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("mysecretkey")

// ValidateToken - Función para validar el token JWT
func ValidateToken(tokenString string) (jwt.Claims, error) {
	// Parsear el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el token esté firmado con el método HS256
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	// Si hay un error en el parseo o el token es inválido
	if err != nil {
		return nil, err
	}

	// Verificar que el token sea válido
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
