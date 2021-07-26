package route

import (
	"awesomeProject/TestTwo/internet/ctrl"
	"github.com/gin-gonic/gin"
)

func ContNum( r *gin.Engine) {
	r.POST("/ping", ctrl.ContController)
}