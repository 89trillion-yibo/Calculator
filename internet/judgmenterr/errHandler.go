package judgmenterr


var (
	OK = response(200,"ok")
	Error = response(500,"服务器错误")

	ParametersCap = response(1001,"参数首部不规范")
	ParametersTail = response(1002,"参数尾部不规范")
	ParametersMiddle = response(1003,"参数中间不规范")
	ParametersIllegal = response(1004,"表达式有非法参数")
	NoData = response(1005,"参数为空")
)

//异常结构
type SoldierErr struct {
	Code    int             //错误码
	Data    interface{}     //返回数据
	Message string          //错误信息
}

//不需要返回数据
func response(code int,message string) *SoldierErr {
	return &SoldierErr{
		Code: code,
		Message: message,
		Data: nil,
	}
}

//需要返回数据
func (sol *SoldierErr) AddData(data interface{}) SoldierErr {
	return SoldierErr{
		Code: sol.Code,
		Message: sol.Message,
		Data: data,
	}
}


