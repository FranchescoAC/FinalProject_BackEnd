package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET_KEY = []byte("mysecretkey")

// JWTMiddleware - Middleware para verificar el token JWT
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "❌ Falta el token de autorización", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWT_SECRET_KEY, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "❌ Token inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
