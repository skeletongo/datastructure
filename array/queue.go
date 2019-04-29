package array

import (
	"bytes"
	"fmt"
)

// 基于动态数组实现的队列
type Queue struct {
	array *Array
}

func (a *Queue) Init() {
	a.array = NewArray()
}

func (a *Queue) GetSize() int {
	return a.array.GetSize()
}

func (a *Queue) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *Queue) Enqueue(e interface{}) {
	a.array.AddLast(e)
}

func (a *Queue) Dequeue() interface{} {
	return a.array.RemoveFirst()
}

func (a *Queue) GetFront() interface{} {
	return a.array.Get(0)
}

func (a *Queue) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("ArrayQueue: front [")
	for i := 0; i < a.array.GetSize(); i++ {
		buf.WriteString(fmt.Sprintf("%v", a.array.Get(i)))
		if i != a.array.GetSize()-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
