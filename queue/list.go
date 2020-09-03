package queue

import (
	"bytes"
	"container/list"
	"fmt"
)

type ListQueue struct {
	*list.List
}

func NewListQueue() *ListQueue {
	return &ListQueue{list.New()}
}

func (l *ListQueue) Enqueue(v interface{}) {
	l.PushBack(v)
}

func (l *ListQueue) Dequeue() interface{} {
	if l.Len() == 0 {
		return nil
	}
	return l.Remove(l.Front())
}

func (l *ListQueue) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d queue: front [", l.Len()))
	if l.Len() > 0 {
		buf.WriteString(fmt.Sprintf("%v", l.Front().Value))
	}
	e := l.Front()
	for i := 1; i < l.Len(); i++ {
		buf.WriteString(fmt.Sprintf(", %v", e.Next().Value))
		e = e.Next()
	}
	buf.WriteString("]")
	return buf.String()
}
