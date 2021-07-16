package route

import (
	"awesomeProject/second/controller"
	"awesomeProject/second/judgmenterr"
	"github.com/gin-gonic/gin"
)

func ContNum( r *gin.Engine) {
	r.POST("/ping", judgmenterr.HandlerErr(controller.ContController))
}