package handler

import "errors"

//判断参数
func Effectiveness(exp string) error{
	if len(exp) == 0 {
		return errors.New("参数为空")
	}
	//判断最后一位是否为操作符
	if exp[len(exp)-1]=='+' || exp[len(exp)-1]=='-' || exp[len(exp)-1]=='*' || exp[len(exp)-1]=='/'{
		return errors.New("参数首尾不规范")
	}
	//判断第一位是否为操作符
	if exp[0]=='+' || exp[0]=='-' || exp[0]=='*' || exp[0]=='/'{
		return errors.New("参数首尾不规范")
	}
	//判断表达式中间是否有相连的操作符
	for i := 1; i < len(exp); i++ {
		if (exp[i]=='+'||exp[i]=='-'||exp[i]=='*'||exp[i]=='/') && (exp[i-1]=='+'||exp[i-1]=='-'||exp[i-1]=='*'||exp[i-1]=='/'){
			return errors.New("参数中间不规范")
		}
	}
	return nil
}
