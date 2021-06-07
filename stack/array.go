package stack

import (
	"bytes"
	"fmt"
)

type ArrayStack struct {
	data []interface{}
}

// NewArrayStack 创建栈
// capacity 初始容量
func NewArrayStack(capacity ...int) *ArrayStack {
	_cap := 0
	if len(capacity) > 0 && capacity[0] > 0 {
		_cap = capacity[0]
	}
	return &ArrayStack{data: make([]interface{}, 0, _cap)}
}

// Len 栈长度
func (a *ArrayStack) Len() int {
	return len(a.data)
}

// Peek 查看栈顶元素
func (a *ArrayStack) Peek() interface{} {
	if a.Len() == 0 {
		return nil
	}
	return a.data[len(a.data)-1]
}

// Push 在栈顶放入元素
func (a *ArrayStack) Push(v interface{}) {
	a.data = append(a.data, v)
}

// Pop 从栈顶取出元素
func (a *ArrayStack) Pop() interface{} {
	if a.Len() == 0 {
		return nil
	}
	index := len(a.data) - 1
	v := a.data[index]
	a.data = a.data[:index]
	return v
}

func (a *ArrayStack) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d stack: [", len(a.data)))
	if len(a.data) > 0 {
		buf.WriteString(fmt.Sprintf("%v", a.data[0]))
	}
	for i := 1; i < len(a.data); i++ {
		buf.WriteString(fmt.Sprintf(", %v", a.data[i]))
	}
	buf.WriteString("] top")
	return buf.String()
}
