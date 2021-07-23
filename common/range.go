package common

import "container/list"

func PreOrder(n BSTNode, f func(interface{})) {
	if IsNil(n) {
		return
	}

	f(n.GetValue())
	PreOrder(n.GetLeftNode(), f)
	PreOrder(n.GetRightNode(), f)
}

func InOrder(n BSTNode, f func(interface{})) {
	if IsNil(n) {
		return
	}

	InOrder(n.GetLeftNode(), f)
	f(n.GetValue())
	InOrder(n.GetRightNode(), f)
}

func PostOrder(n BSTNode, f func(interface{})) {
	if IsNil(n) {
		return
	}

	PostOrder(n.GetLeftNode(), f)
	PostOrder(n.GetRightNode(), f)
	f(n.GetValue())
}

func LevelOrder(n BSTNode, f func(interface{})) {
	if IsNil(n) {
		return
	}

	q := list.New()
	q.PushBack(n)
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		node := e.Value.(BSTNode)
		f(node.GetValue())
		if !IsNil(node.GetLeftNode()) {
			q.PushBack(node.GetLeftNode())
		}
		if !IsNil(node.GetRightNode()) {
			q.PushBack(node.GetRightNode())
		}
	}
}
