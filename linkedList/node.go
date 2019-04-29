package linkedList

// =====================
// 单向链表节点
// =====================
type node struct {
	e    interface{}
	next *node
}

func newNode(e interface{}, next *node) *node {
	return &node{e, next}
}

// =====================
// 双向链表节点
// =====================
type node2 struct {
	e    interface{}
	prev *node2
	next *node2
}

func newNode2(e interface{}, prev, next *node2) *node2 {
	return &node2{e, prev, next}
}
