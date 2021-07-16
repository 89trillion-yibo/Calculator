package handler

import (
	"awesomeProject/TestTwo/util"
	"fmt"
	"strconv"
)
//中缀表达式转后缀表达式
func InfixToPostfix(infix string) []string {
	//初始化栈
	stack := util.Stack{}
	postfix := []string{}
	len := len(infix)

	//遍历字符串
	for i := 0; i < len; i++ {
		char := string(infix[i])
		if char == ""{
			continue
		}else if char >= "0" && char <= "9"{
			k := i
			tmp := ""
			for ; k < len && infix[k] >= '0' && infix[k] <= '9'; k++ {
				tmp += string(infix[k])
			}
			postfix = append(postfix, tmp)
			i = k-1
		}else if char == "+" || char == "-" || char == "*" || char == "/"{
			for !stack.IsEmpty() {
				peek, err := stack.Peek()
				if err!=nil{
					fmt.Println(err)
				}
				//比较
				if  compare(peek.(string),char){
					break
				}
				postfix = append(postfix,peek.(string))
				stack.Pop()
			}
			//操作符入栈
			stack.Push(char)
		}
	}

	for !stack.IsEmpty(){
		pop, _ := stack.Pop()
		postfix = append(postfix, pop.(string))
	}

	return postfix
}

func compare(peek string, newTop string) bool {
	if peek == "+" || peek == "-" {
		if newTop == "*" || newTop == "/"{
			return true
		}
	}
	return false
}

//后缀表达式计算值
func Calculate(postfix []string) int {
	stack := util.Stack{}
	len := len(postfix)

	for i := 0; i < len; i++ {
		//数字直接压栈
		if postfix[i] >= "0" && postfix[i] <= "9"{
			stack.Push(postfix[i])
		}else {
			//操作符:取出两个数字计算值，再将结果压栈
			pop1, _ := stack.Pop()
			pop2, _ := stack.Pop()
			num1, _ := strconv.Atoi(pop1.(string))
			num2, _ := strconv.Atoi(pop2.(string))
			switch postfix[i] {
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num2 - num1))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "/":
				stack.Push(strconv.Itoa(num2 / num1))
			}
		}
	}
	peek, _ := stack.Peek()
	result, _ := strconv.Atoi(peek.(string))
	return result
}
