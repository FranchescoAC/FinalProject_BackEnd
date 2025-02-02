package main

import (
	"log"
	"net/http"
	"userManagement/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Crear el router
	r := mux.NewRouter()

	// Configurar las rutas
	routes.SetRoutes(r)

	// Iniciar el servidor en el puerto 2003
	log.Println("Servidor corriendo en http://localhost:2003")
	log.Fatal(http.ListenAndServe(":2003", r))
}
