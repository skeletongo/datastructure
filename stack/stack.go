// 栈
package stack

type Stack interface {
	Len() int
	Peek() interface{}
	Push(interface{})
	Pop() interface{}
}
