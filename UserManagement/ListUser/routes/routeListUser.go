package routes

import (
	"userManagement/controller"
	"userManagement/middleware"
	"net/http"
	"github.com/gorilla/mux"
)

// SetRoutes - Set up the routes for the ListUser service
func SetRoutes(r *mux.Router) {
	r.Handle("/users", middleware.JWTMiddleware(http.HandlerFunc(controller.GetUsers))).Methods("GET")
}
