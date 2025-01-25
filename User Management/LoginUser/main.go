package main

import (
	"github.com/gin-gonic/gin"
	"LoginUser/config"
	"LoginUser/routes"
)

func main() {
	config.InitDB()
	defer config.GetDB().Close()

	r := gin.Default()
	routes.LoginRoutes(r)

	r.Run(":8082")
}