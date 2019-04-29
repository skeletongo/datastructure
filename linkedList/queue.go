package linkedList

import (
	"bytes"
	"fmt"
)

// 基于单向链表实现的队列
type Queue struct {
	head *node // 头节点
	tail *node // 尾节点,使在尾部添加节点的时间复杂度为O(1)
	size int
}

func (l *Queue) GetSize() int {
	return l.size
}

func (l *Queue) IsEmpty() bool {
	return l.size == 0
}

func (l *Queue) Enqueue(e interface{}) {
	if l.tail == nil {
		l.tail = newNode(e, nil)
		l.head = l.tail
	} else {
		l.tail.next = newNode(e, nil)
		l.tail = l.tail.next
	}
	l.size++
}

func (l *Queue) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("no data")
	}
	res := l.head
	l.head = l.head.next
	res.next = nil

	// 最后一个元素出队后尾节点也要置空
	if l.head == nil {
		l.tail = nil
	}

	l.size--
	return res.e
}

func (l *Queue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("no data")
	}
	return l.head.e
}

func (l *Queue) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("LinkedListQueue: front [")
	cur := l.head
	for cur != nil {
		buf.WriteString(fmt.Sprintf("%v", cur.e))
		if cur.next != nil {
			buf.WriteString(",")
		}
		cur = cur.next
	}
	buf.WriteString("]")
	return buf.String()
}
