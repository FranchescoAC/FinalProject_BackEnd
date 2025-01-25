package routes

import (
	"net/http"
	"registerUser/controller"
)

func RegisterRoutes() {
	http.HandleFunc("/register", controller.RegisterUser)
}
