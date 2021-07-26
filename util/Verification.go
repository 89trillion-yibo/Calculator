package util

//判断参数
func Effectiveness(exp string) string{
	//判断最后一位是否为操作符
	if exp[len(exp)-1]=='+' || exp[len(exp)-1]=='-' || exp[len(exp)-1]=='*' || exp[len(exp)-1]=='/'{
		return "1002"
	}
	//判断第一位是否为操作符
	if exp[0]=='+' || exp[0]=='-' || exp[0]=='*' || exp[0]=='/'{
		return "1001"
	}
	//判断表达式中间是否有相连的操作符
	for i := 1; i < len(exp); i++ {
		if (exp[i]=='+'||exp[i]=='-'||exp[i]=='*'||exp[i]=='/') && (exp[i-1]=='+'||exp[i-1]=='-'||exp[i-1]=='*'||exp[i-1]=='/'){
			return "1003"
		}
	}
	//判断是否有除0-9和+,-,*,/外的非法字符
	for i := 0; i < len(exp); i++ {
		if (exp[i]<'0' || exp[i]>'9') && (exp[i]!='+' && exp[i]!='-' && exp[i]!='*' && exp[i]!='/') {
			return "1004"
		}
	}
	return ""
}
