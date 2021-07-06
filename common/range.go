package common

import "container/list"

func PreOrder(n INode, f func(INode)) {
	if IsNil(n) {
		return
	}

	f(n)
	PreOrder(n.GetLeftNode(), f)
	PreOrder(n.GetRightNode(), f)
}

func InOrder(n INode, f func(INode)) {
	if IsNil(n) {
		return
	}

	InOrder(n.GetLeftNode(), f)
	f(n)
	InOrder(n.GetRightNode(), f)
}

func PostOrder(n INode, f func(INode)) {
	if IsNil(n) {
		return
	}

	PostOrder(n.GetLeftNode(), f)
	PostOrder(n.GetRightNode(), f)
	f(n)
}

func LevelOrder(n INode, f func(node INode)) {
	if IsNil(n) {
		return
	}

	q := list.New()
	q.PushBack(n)
	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		node := e.Value.(INode)
		f(node)
		if !IsNil(node.GetLeftNode()) {
			q.PushBack(node.GetLeftNode())
		}
		if !IsNil(node.GetRightNode()) {
			q.PushBack(node.GetRightNode())
		}
	}
}
