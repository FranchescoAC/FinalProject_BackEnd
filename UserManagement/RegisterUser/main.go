package main

import (
	"log"
	"net/http"
	"userManagement/routes"
)

func main() {
	// Registrar las rutas que hemos definido
	r := routes.RegisterRoutes()

	// Configurar el puerto del servidor
	port := ":2000" // Puerto 2000 para el microservicio RegisterUser

	// Iniciar el servidor HTTP
	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, r))  // El servidor escucha en el puerto 2000
}

