//
// 时间复杂度：
package ring

import (
	"bytes"
	"fmt"
)

// Ring 循环队列
type Ring struct {
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

// New 创建队列
// capacity 初始容量
func New(capacity ...int) *Ring {
	// 初始容量默认为1
	// 因为扩容是在原来的容量上扩大两倍
	// 实际数组长度要比容量大1
	// 这是为了区分对列为空和满的状态
	// 如果数组长度和容量相等，则队列为空和队列已满时 q.front == q.tail 都为真
	_cap := 2
	if len(capacity) > 0 && capacity[0] > 0 {
		_cap = capacity[0] + 1
	}
	return &Ring{data: make([]interface{}, _cap)}
}

// Len 队列长度
func (q *Ring) Len() int {
	return q.len
}

// Cap 队列容量
func (q *Ring) Cap() int {
	return len(q.data) - 1
}

// IsEmpty 是否为空队列
func (q *Ring) IsEmpty() bool {
	return q.front == q.tail
}

// resize 容量调整
func (q *Ring) resize(n int) {
	newData := make([]interface{}, n+1)
	for i, j := q.front, 0; j < q.len; j++ {
		newData[j] = q.data[i]
		i++
		i %= len(q.data)
	}
	q.data = newData
	q.front = 0
	q.tail = q.len
}

// Enqueue 放入一个元素
func (q *Ring) Enqueue(e interface{}) {
	if q.len == q.Cap() {
		q.resize(q.len * 2)
	}
	q.data[q.tail] = e
	q.tail++
	q.tail %= len(q.data)
	q.len++
}

// Dequeue 取出一个元素
func (q *Ring) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	if q.len == q.Cap()/4 && q.Cap()/2 != 0 {
		q.resize(q.Cap() / 2)
	}
	e := q.data[q.front]
	q.data[q.front] = nil
	q.front++
	q.front %= len(q.data)
	q.len--
	return e
}

// Front 查看头部元素
func (q *Ring) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.data[q.front]
}

func (q *Ring) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d queue: front [", q.len))
	if q.len > 0 {
		buf.WriteString(fmt.Sprintf("%v", q.data[q.front]))
	}
	for i, j := (q.front+1)%q.Cap(), 0; j < q.len-1; j++ {
		buf.WriteString(fmt.Sprintf(", %v", q.data[i]))
		i++
		i %= len(q.data)
	}
	buf.WriteString("]")
	return buf.String()
}
