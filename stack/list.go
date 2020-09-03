package stack

import (
	"bytes"
	"container/list"
	"fmt"
)

type ListStack struct {
	*list.List
}

func NewListStack() *ListStack {
	return &ListStack{list.New()}
}

func (l *ListStack) Peek() interface{} {
	if l.Len() == 0 {
		return nil
	}
	return l.Back().Value
}

func (l *ListStack) Push(v interface{}) {
	l.PushBack(v)
}

func (l *ListStack) Pop() interface{} {
	if l.Len() == 0 {
		return nil
	}
	return l.Remove(l.Back())
}

func (l *ListStack) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d stack: [", l.Len()))
	if l.Len() > 0 {
		buf.WriteString(fmt.Sprintf("%v", l.Front().Value))
	}
	e := l.Front()
	for i := 1; i < l.Len(); i++ {
		buf.WriteString(fmt.Sprintf(", %v", e.Next().Value))
		e = e.Next()
	}
	buf.WriteString("] top")
	return buf.String()
}
