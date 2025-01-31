package main

import (
	"github.com/gin-gonic/gin"
	"UpdateUser/config"
	"UpdateUser/routes"
)

func main() {
	config.InitDB()
	defer config.GetDB().Close()

	r := gin.Default()
	routes.UpdateRoutes(r)

	r.Run(":8084")
}
