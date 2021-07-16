package main

import (
	"awesomeProject/second/service"
	"testing"
)

func TestConversion(t *testing.T)  {
	exp := "1+2*4-6/3"
	if !service.Effectiveness(exp){
		t.Log("输入错误")
	}
	postfix := service.InfixToPostfix(exp)
	calculate := service.Calculate(postfix)
	t.Log(calculate)
}