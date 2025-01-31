package routes

import (
	"github.com/gin-gonic/gin"
	"UpdateUser/controller"
)

func UpdateRoutes(router *gin.Engine) {
	controller := controller.UpdateController{}
	router.PUT("/update/:id", controller.UpdateUser)
}