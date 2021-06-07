// 基于循环链表实现的队列，增加和取出元素的时间复杂度为O(1)
package queue

import (
	"bytes"
	"container/ring"
	"fmt"
)

type RingQueue struct {
	*ring.Ring
	len int
}

func NewRingQueue() *RingQueue {
	return new(RingQueue)
}

func (r *RingQueue) Len() int {
	return r.len
}

func (r *RingQueue) Enqueue(v interface{}) {
	r.len++
	if r.Ring == nil {
		r.Ring = ring.New(1)
		r.Value = v
		return
	}
	e := ring.New(1)
	e.Value = v
	r.Ring = r.Prev().Link(e)
}

func (r *RingQueue) Dequeue() interface{} {
	if r.len == 0 {
		return nil
	}
	r.len--
	if r.len == 0 {
		v := r.Value
		r.Ring = nil
		return v
	}
	r.Ring = r.Prev()
	v := r.Unlink(1).Value
	r.Ring = r.Next()
	return v
}

func (r *RingQueue) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d queue: front [", r.len))
	if r.len > 0 {
		buf.WriteString(fmt.Sprintf("%v", r.Value))
	}
	for i := 1; i < r.len; i++ {
		r.Ring = r.Move(1)
		buf.WriteString(fmt.Sprintf(", %v", r.Value))
	}
	r.Ring = r.Move(1)
	buf.WriteString("]")
	return buf.String()
}
