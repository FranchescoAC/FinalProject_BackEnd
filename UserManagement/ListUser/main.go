package main

import (
	"log"
	"net/http"
	"userManagement/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Crear un router
	r := mux.NewRouter()

	// Configurar las rutas
	routes.SetRoutes(r)

	// Iniciar el servidor en el puerto 2002
	log.Println("Servidor corriendo en http://localhost:2002")
	log.Fatal(http.ListenAndServe(":2002", r))
}
