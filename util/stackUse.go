package util

import "errors"

//声明stack为空接口类型的切片
type Stack []interface{}

//栈长度
func (stack Stack) Len() int{
	return len(stack)
}
//栈是否为空
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
//压栈
func (stack *Stack) Push(val interface{})  {
	*stack = append(*stack,val)
}
//栈顶部数据
func (stack Stack) Peek() (interface{}, error) {
	if len(stack) == 0 {
		return nil,errors.New("Out of index,len is 0")
	}
	return stack[len(stack)-1],nil
}
//出栈
func (stack *Stack) Pop() (interface{},error)  {
	theStack := *stack
	if len(*stack) == 0 {
		return nil,errors.New("Out of index,len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value,nil
}

