package ctrl

import (
	"awesomeProject/TestTwo/internet/handler"
	"github.com/gin-gonic/gin"
)

func ContController(c *gin.Context) (interface{},error) {
	exp := c.PostForm("exp")
	err := handler.Effectiveness(exp)
    if err != nil{
    	return nil,err
	}
	postfix := handler.InfixToPostfix(exp)
	calculate := handler.Calculate(postfix)
	return calculate,nil
}
