package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"registerUser/controller"
	"github.com/rs/cors"
)

func RegisterRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterUser).Methods("POST")

	// Enable CORS for all routes
	handler := cors.Default().Handler(r)

	// Start the server
	http.Handle("/", handler)
	http.ListenAndServe(":8081", nil)
}
