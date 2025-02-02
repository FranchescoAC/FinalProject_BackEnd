package controller

import (
	"encoding/json"
	"net/http"
	"userManagement/model"
	"github.com/golang-jwt/jwt/v4"  // Asegúrate de importar jwt/v4
	"time"
)

// Función para hacer login de un usuario
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	// Decodificar la solicitud JSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Verificar que el usuario exista en la base de datos
	storedUser, err := model.GetUserByEmail(user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Verificar que las contraseñas coincidan
	if storedUser.Password != user.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Crear el token JWT
	token := jwt.New(jwt.SigningMethodHS256)  // Esta línea ahora funciona bien en la versión 4.x

	// Configuración de los claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = storedUser.ID
	claims["username"] = storedUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Firmar el token
	tokenString, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Responder con el token JWT
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
