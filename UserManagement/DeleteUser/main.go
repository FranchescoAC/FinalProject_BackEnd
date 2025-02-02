package main

import (
	"github.com/gin-gonic/gin"
	"DeleteUser/config"
	"DeleteUser/routes"
)

func main() {
	config.InitDB()
	defer config.GetDB().Close()

	r := gin.Default()
	routes.DeleteRoutes(r)

	r.Run(":8085")
}