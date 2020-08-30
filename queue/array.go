//
// 时间复杂度：
package queue

import (
	"bytes"
	"fmt"
)

// ArrayQueue 数组循环队列
type ArrayQueue struct {
	// 数组,将切片当数组用
	// 因为go语言的数组元素个数也是数组类型的一部分
	// 所以这里不能用数组声明，否则不能扩容
	// 这里为了演示用slice代替
	data []interface{}

	// 第一个元素位置
	front int

	// 下一个新增元素的位置
	tail int

	// 元素数量
	len int
}

// NewArrayQueue 创建队列
// capacity 初始容量
func NewArrayQueue(capacity ...int) *ArrayQueue {
	// 初始容量默认为1
	// 因为扩容是在原来的容量上扩大两倍
	// 实际数组长度要比容量大1
	// 这是为了区分对列为空和满的状态
	// 如果数组长度和容量相等，则队列为空和队列已满时 q.front == q.tail 都为真
	_cap := 2
	if len(capacity) > 0 && capacity[0] > 0 {
		_cap = capacity[0] + 1
	}
	return &ArrayQueue{data: make([]interface{}, _cap)}
}

// Len 队列长度
func (a *ArrayQueue) Len() int {
	return a.len
}

// Cap 队列容量
func (a *ArrayQueue) Cap() int {
	return len(a.data) - 1
}

// resize 容量调整
func (a *ArrayQueue) resize(n int) {
	newData := make([]interface{}, n+1)
	for i, j := a.front, 0; j < a.len; j++ {
		newData[j] = a.data[i]
		i++
		i %= len(a.data)
	}
	a.data = newData
	a.front = 0
	a.tail = a.len
}

// Enqueue 放入一个元素
func (a *ArrayQueue) Enqueue(v interface{}) {
	if a.len == a.Cap() {
		a.resize(a.len * 2)
	}
	a.data[a.tail] = v
	a.tail++
	a.tail %= len(a.data)
	a.len++
}

// Dequeue 取出一个元素
func (a *ArrayQueue) Dequeue() interface{} {
	if a.len == 0 {
		return nil
	}
	if a.len == a.Cap()/4 && a.Cap()/2 != 0 {
		a.resize(a.Cap() / 2)
	}
	e := a.data[a.front]
	a.data[a.front] = nil
	a.front++
	a.front %= len(a.data)
	a.len--
	return e
}

// Front 查看头部元素
func (a *ArrayQueue) Front() interface{} {
	if a.len == 0 {
		return nil
	}
	return a.data[a.front]
}

func (a *ArrayQueue) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d queue: front [", a.len))
	if a.len > 0 {
		buf.WriteString(fmt.Sprintf("%v", a.data[a.front]))
	}
	for i, j := (a.front+1)%a.Cap(), 0; j < a.len-1; j++ {
		buf.WriteString(fmt.Sprintf(", %v", a.data[i]))
		i++
		i %= len(a.data)
	}
	buf.WriteString("]")
	return buf.String()
}
