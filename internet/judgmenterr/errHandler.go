package judgmenterr

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NumHandler func(c *gin.Context) (interface{}, error)

func HandlerErr (contNum NumHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := contNum(c)
		if err!=nil{
			contError := err.(ContError)
			c.JSON(contError.Status,contError)
		}
		c.JSON(http.StatusOK,gin.H{
			"计算结果" : data,
		})
	}
}