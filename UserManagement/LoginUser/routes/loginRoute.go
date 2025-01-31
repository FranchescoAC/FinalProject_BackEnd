package routes

import (
	"github.com/gin-gonic/gin"
	"LoginUser/controller"
)

func LoginRoutes(router *gin.Engine) {
	controller := controller.LoginController{}
	router.POST("/login", controller.LoginUser)
}


