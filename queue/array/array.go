//
// 时间复杂度：
package array

import (
	"bytes"
	"fmt"
)

// Array 队列
type Array struct {
	data []interface{}
}

// New 创建队列
// capacity 初始容量
func New(capacity ...int) *Array {
	_cap := 0
	if len(capacity) > 0 && capacity[0] > 0 {
		_cap = capacity[0]
	}
	return &Array{data: make([]interface{}, 0, _cap)}
}

// Len 队列长度
func (q *Array) Len() int {
	return len(q.data)
}

// Cap 队列容量
func (q *Array) Cap() int {
	return cap(q.data)
}

// IsEmpty 是否为空队列
func (q *Array) IsEmpty() bool {
	return len(q.data) == 0
}

// Enqueue 放入一个元素
func (q *Array) Enqueue(e interface{}) {
	q.data = append(q.data, e)
}

// Dequeue 取出一个元素
func (q *Array) Dequeue() interface{} {
	if len(q.data) <= 0 {
		return nil
	}
	e := q.data[0]
	copy(q.data, q.data[1:])
	q.data = q.data[:len(q.data)-1]
	return e
}

// Front 查看头部元素
func (q *Array) Front() interface{} {
	if len(q.data) <= 0 {
		return nil
	}
	return q.data[0]
}

func (q *Array) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d queue: front [", len(q.data)))
	if len(q.data) > 0 {
		buf.WriteString(fmt.Sprintf("%v", q.data[0]))
	}
	for i := 1; i < len(q.data); i++ {
		buf.WriteString(fmt.Sprintf(", %v", q.data[i]))
	}
	buf.WriteString("]")
	return buf.String()
}
