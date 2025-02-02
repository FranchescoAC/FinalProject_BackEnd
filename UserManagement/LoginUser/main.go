package main

import (
	"log"
	"net/http"
	"userManagement/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Crear un nuevo enrutador
	r := mux.NewRouter()

	// Registrar las rutas de LoginUser
	routes.RegisterRoutes(r)

	// Iniciar el servidor en el puerto 2001
	log.Fatal(http.ListenAndServe(":2001", r))
}

