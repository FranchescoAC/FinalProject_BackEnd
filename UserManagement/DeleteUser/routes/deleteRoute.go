package routes

import (
	"github.com/gin-gonic/gin"
	"DeleteUser/controller"
)

func DeleteRoutes(router *gin.Engine) {
	controller := controller.DeleteController{}
	router.DELETE("/delete/:id", controller.DeleteUser)
}