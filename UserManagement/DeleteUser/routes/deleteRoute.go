package routes

import (
	"userManagement/controller"
	"userManagement/middleware"
	"net/http"
	"github.com/gorilla/mux"
)

// SetRoutes - Configure routes for the DeleteUser service
func SetRoutes(r *mux.Router) {
	r.Handle("/deleteUser", middleware.JWTMiddleware(http.HandlerFunc(controller.DeleteUser))).Methods("DELETE")
}
