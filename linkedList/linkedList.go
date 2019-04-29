package linkedList

import (
	"bytes"
	"fmt"
	"reflect"
)

// 单向链表
// 特点：在头部添加或删除元素时间复杂度是O(1),在尾部追加或删除元素时间复杂度是O(n)
type LinkedList struct {
	dummyHead node
	size      int
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在指定位置添加新节点
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("invalid index")
	}
	prev := &l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = newNode(e, prev.next)
	l.size++
}

func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size { // l.size 是待添加元素，所以不包含l.size
		panic("invalid index")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size { // l.size 是待添加元素，所以不包含l.size
		panic("invalid index")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for cur != nil {
		if reflect.DeepEqual(cur.e, e) {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size { // l.size 是待添加元素，所以不包含l.size
		panic("invalid index")
	}
	prev := &l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	res := prev.next
	prev.next = res.next
	res.next = nil
	l.size--
	return res.e
}

func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) RemoveElement(e interface{}) {
	//prev := &l.dummyHead
	//for i := 0; i < l.size; i++ {
	//	if reflect.DeepEqual(prev.next.e, e) {
	//		delNode := prev.next
	//		prev.next = delNode.next
	//		delNode.next = nil
	//		l.size--
	//		return
	//	}
	//	prev = prev.next
	//}

	prev := &l.dummyHead
	for prev.next != nil { // 有节点就继续循环
		if reflect.DeepEqual(prev.next.e, e) {
			break
		}
		prev = prev.next
	}

	if prev.next != nil { // 不为nil说明找到了相同元素
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
		l.size--
	}
}

func (l *LinkedList) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("LinkedList: front [")
	cur := l.dummyHead.next
	for cur != nil {
		buf.WriteString(fmt.Sprintf("%v", cur.e))
		if cur.next != nil {
			buf.WriteString(",")
		}
		cur = cur.next
	}
	buf.WriteString("] end")
	return buf.String()
}
