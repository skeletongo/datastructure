package array

import (
	"bytes"
	"fmt"
)

// 基于数组实现的循环队列(动态数组)
type LoopQueue struct {
	array []interface{}
	front int
	tail  int
	size  int
}

func (l *LoopQueue) Init() {
	l.array = make([]interface{}, DefaultCapacity)
}

func (l *LoopQueue) GetSize() int {
	return l.size
}

func (l *LoopQueue) IsEmpty() bool {
	return l.front == l.tail
}

// 队列容量
func (l *LoopQueue) GetCapacity() int {
	return len(l.array) - 1
}

// 修改队列容量
// num 队列容量（实际数组长度要比容量多一个）
func (l *LoopQueue) resize(num int) {
	data := make([]interface{}, num+1)
	for i := 0; i < l.size; i++ {
		data[i] = l.array[(i+l.front)%len(l.array)]
	}
	l.array = data
	l.front = 0
	l.tail = l.size
}

func (l *LoopQueue) Enqueue(e interface{}) {
	// 检查扩容
	if (l.tail+1)%len(l.array) == l.front {
		l.resize(l.GetCapacity() * 2)
	}
	l.array[l.tail] = e
	// 维护tail
	l.tail = (l.tail + 1) % len(l.array)
	l.size++
}

func (l *LoopQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("no data")
	}
	res := l.array[l.front]
	l.array[l.front] = nil
	l.front = (l.front + 1) % len(l.array)
	l.size--
	// 检查缩容
	if l.size == l.GetCapacity()/4 && l.GetCapacity()/2 != 0 {
		l.resize(l.GetCapacity() / 2)
	}
	return res
}

func (l *LoopQueue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("no data")
	}
	return l.array[l.front]
}

func (l *LoopQueue) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("LoopQueue: front [")
	for i := l.front; i != l.tail; i = (i + 1) % len(l.array) {
		buf.WriteString(fmt.Sprintf("%v", l.array[i]))
		if (i+1)%len(l.array) != l.tail {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
