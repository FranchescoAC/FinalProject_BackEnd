package controller

import (
	"encoding/json" // Asegúrate de importar este paquete para poder usar JSON
	"net/http"
	"userManagement/model" // Asegúrate de importar el modelo
	"github.com/dgrijalva/jwt-go" // Para el manejo de JWT
	"time"
)

var JWT_SECRET_KEY = []byte("your_secret_key") // Clave secreta para la creación de JWT

// Registrar un nuevo usuario
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	// Decodificamos el JSON del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Llamamos a la función que crea el usuario en la base de datos
	createdUser, err := model.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Respondemos con un código 201 y los datos del nuevo usuario en formato JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// Login del usuario (genera un JWT)
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decodificamos las credenciales
	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Lógica para verificar las credenciales (simulada)
	if loginDetails.Username == "test" && loginDetails.Password == "password" { // Reemplaza esto con la lógica real de validación
		tokenString, err := generateJWT(loginDetails.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Respondemos con el JWT generado
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

// Función para generar el token JWT
func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token válido por 1 día
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
