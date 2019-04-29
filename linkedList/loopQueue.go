package linkedList

import (
	"bytes"
	"fmt"
)

// 双向循环队列
type LoopQueue struct {
	dummyHead node2
	size      int
}

func (l *LoopQueue) GetSize() int {
	return l.size
}

func (l *LoopQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *LoopQueue) Enqueue(e interface{}) {
	if l.dummyHead.next == nil {
		l.dummyHead.next = newNode2(e, &l.dummyHead, &l.dummyHead)
		l.dummyHead.prev = l.dummyHead.next
	} else {
		prev := l.dummyHead.prev
		l.dummyHead.prev = newNode2(e, l.dummyHead.prev, &l.dummyHead)
		prev.next = l.dummyHead.prev
	}
	l.size++
}

func (l *LoopQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("no data")
	}
	res := l.dummyHead.next
	if res.next == &l.dummyHead {
		l.dummyHead.next = nil
		l.dummyHead.prev = nil
	} else {
		l.dummyHead.next = res.next
		res.prev = &l.dummyHead
	}

	res.next = nil
	res.prev = nil

	l.size--
	return res.e
}

func (l *LoopQueue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("no data")
	}
	return l.dummyHead.next.e
}

func (l *LoopQueue) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("LoopQueue: front [")
	cur := l.dummyHead.next
	for cur != &l.dummyHead {
		buf.WriteString(fmt.Sprintf("%v", cur.e))
		if cur.next != &l.dummyHead {
			buf.WriteString(",")
		}
		cur = cur.next
	}
	buf.WriteString("]")
	return buf.String()
}
