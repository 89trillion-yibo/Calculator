package main

import (
	"awesomeProject/second/route"
	"github.com/gin-gonic/gin"
)

func StartRun() {
	r := gin.Default()
	route.ContNum(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务器
}
