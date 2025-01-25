package main

import (
	"fmt"
	"net/http"
	"registerUser/config"
	"registerUser/routes"
)

func main() {
	config.ConnectDatabase()
	config.InitDatabase()

	routes.RegisterRoutes()

	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", nil)
}