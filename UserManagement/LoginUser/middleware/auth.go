package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET_KEY = []byte("mysecretkey")

// Middleware para verificar el token y el rol del usuario
func JWTMiddleware(next http.Handler, requiredRole string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener el token del encabezado Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Extraer el token (Formato esperado: "Bearer <token>")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWT_SECRET_KEY, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Obtener los claims del token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Verificar el rol del usuario
		userRole := claims["role"].(string)
		if requiredRole != "" && userRole != requiredRole {
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		// Pasar al siguiente middleware si es válido
		next.ServeHTTP(w, r)
	})
}
