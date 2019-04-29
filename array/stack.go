package array

import (
	"bytes"
	"fmt"
)

// 基于动态数组实现的栈
type Stack struct {
	array *Array
}

func (a *Stack) Init() {
	a.array = NewArray()
}

func (a *Stack) GetSize() int {
	return a.array.GetSize()
}

func (a *Stack) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *Stack) Push(e interface{}) {
	a.array.AddLast(e)
}

func (a *Stack) Pop() interface{} {
	return a.array.RemoveLast()
}

func (a *Stack) Peek() interface{} {
	return a.array.Get(a.array.GetSize() - 1)
}

func (a *Stack) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("ArrayStack: [")
	for i := 0; i < a.array.GetSize(); i++ {
		buf.WriteString(fmt.Sprintf("%v", a.array.Get(i)))
		if i != a.array.GetSize()-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("] top")
	return buf.String()
}
