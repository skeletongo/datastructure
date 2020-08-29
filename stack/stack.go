//
// 时间复杂度：
package stack

import (
	"bytes"
	"fmt"
)

// Stack 栈
type Stack struct {
	data []interface{}
}

// New 创建栈
// capacity 初始容量
func New(capacity ...int) *Stack {
	_cap := 0
	if len(capacity) > 0 && capacity[0] > 0 {
		_cap = capacity[0]
	}
	return &Stack{data: make([]interface{}, 0, _cap)}
}

// Len 栈长度
func (s *Stack) Len() int {
	return len(s.data)
}

// Cap 容量
func (s *Stack) Cap() int {
	return cap(s.data)
}

// IsEmpty 栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Push 在栈顶放入元素
func (s *Stack) Push(e interface{}) {
	s.data = append(s.data, e)
}

// Pop 从栈顶取出元素
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	index := len(s.data) - 1
	e := s.data[index]
	s.data = s.data[:index]
	return e
}

// Peek 查看栈顶元素
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d stack: [", len(s.data)))
	if len(s.data) > 0 {
		buf.WriteString(fmt.Sprintf("%v", s.data[0]))
	}
	for i := 1; i < len(s.data); i++ {
		buf.WriteString(fmt.Sprintf(", %v", s.data[i]))
	}
	buf.WriteString("] top")
	return buf.String()
}
