package utils

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

// Middleware para validar el JWT
func ValidateJWT(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	// Obtener el token de las cabeceras de la solicitud
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return nil, fmt.Errorf("Authorization header is missing")
	}

	// Eliminar el prefijo 'Bearer ' del token
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Verifica el token utilizando la clave secreta
	secretKey := []byte("your-secret-key") // Usa una clave secreta de entorno
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Asegurarse de que el algoritmo sea HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return nil, err
	}

	return token, nil
}