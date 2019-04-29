package linkedList

import (
	"bytes"
	"fmt"
)

// 基于单向链表实现的栈
type Stack struct {
	list LinkedList
}

func (l *Stack) GetSize() int {
	return l.list.GetSize()
}

func (l *Stack) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *Stack) Push(e interface{}) {
	l.list.AddFirst(e)
}

func (l *Stack) Pop() interface{} {
	return l.list.RemoveFirst()
}

func (l *Stack) Peek() interface{} {
	return l.list.GetFirst()
}

func (l *Stack) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("LinkedListStack: top [")
	for i := 0; i < l.list.GetSize(); i++ {
		buf.WriteString(fmt.Sprintf("%v", l.list.Get(i)))
		if i != l.list.GetSize()-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
