package main

import (
	"github.com/gin-gonic/gin"
	"ViewUser/config"
	"ViewUser/routes"
)

func main() {
	config.InitDB()
	defer config.GetDB().Close()

	r := gin.Default()
	routes.ViewRoutes(r)

	r.Run(":8083")
}
