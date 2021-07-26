package ctrl

import (
	"awesomeProject/TestTwo/internet/handler"
	"awesomeProject/TestTwo/internet/judgmenterr"
	"awesomeProject/TestTwo/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ContController(c *gin.Context) {
	exp,a := c.GetPostForm("exp")
	fmt.Println("exp:"+exp)
	if !a || exp == "" {
		c.JSON(http.StatusBadRequest,judgmenterr.NoData)
	}else {
		//判断参数正确性
		code := util.Effectiveness(exp)
		if code != ""{
			switch code {
			case "1001":
				c.JSON(http.StatusBadRequest,judgmenterr.ParametersCap)
			case "1002":
				c.JSON(http.StatusBadRequest,judgmenterr.ParametersTail)
			case "1003":
				c.JSON(http.StatusBadRequest,judgmenterr.ParametersMiddle)
			case "1004":
				c.JSON(http.StatusBadRequest,judgmenterr.ParametersIllegal)
			}
		}else {
			postfix := handler.InfixToPostfix(exp)
			calculate := handler.Calculate(postfix)
			c.JSON(http.StatusOK,judgmenterr.OK.AddData(calculate))
		}
	}
}
