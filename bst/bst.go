//二分搜索树的查询,添加,删除操作的时间复杂度为O(h) 最差时间复杂度O(n)链表 最佳时间复杂O(logN)满二叉树
//
//满二叉树：除了最大层节点，其他节点都有两个子节点
//完全二叉树：按每层从左到右的位置添加新节点
//平衡二叉树：所有叶子节点所在的层数的差值的绝对值不能大于1
// 时间复杂度：
package bst

import (
	"dataStructure/queue"
	"dataStructure/stack"
)

type Node struct {
	left, right *Node
	Value       interface{}
}

// BST 二分搜索树
type BST struct {
	root    *Node
	size    int
	Compare func(a, b interface{}) int
}

// 创建二分搜索树
// 参数 f 为自定义元素大小比较函数
// 大小比较函数 返回值：
// -1	表示	a<b
// 0	表示	a=b
// 1	表示	a>b
func NewBST(f func(a, b interface{}) int) *BST {
	return &BST{Compare: f}
}

// GetSize 获取节点总数
func (b *BST) GetSize() int {
	return b.size
}

// IsEmpty 是否为空树
func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// Add 添加新节点
func (b *BST) Add(v interface{}) {
	b.root = b.add(b.root, v)
}

func (b *BST) add(node *Node, v interface{}) *Node {
	if node == nil {
		b.size++
		return &Node{Value: v}
	}

	switch b.Compare(v, node.Value) {
	case -1:
		node.left = b.add(node.left, v)
	case 1:
		node.right = b.add(node.right, v)
	}
	return node
}

// Add 添加新节点非递归
func (b *BST) Add2(v interface{}) {
	if b.size == 0 {
		b.root = &Node{Value: v}
		return
	}

	node := b.root
	for {
		switch b.Compare(node.Value, v) {
		case 0:
			return
		case -1:
			if node.right == nil {
				b.size++
				node.right = &Node{Value: v}
				return
			}
			node = node.right
		case 1:
			if node.left == nil {
				b.size++
				node.left = &Node{Value: v}
				return
			}
			node = node.left
		}
	}
}

// Contains 查询是否包含指定元素
func (b *BST) Contains(v interface{}) bool {
	return b.contains(b.root, v)
}

func (b *BST) contains(node *Node, v interface{}) bool {
	if node == nil {
		return false
	}

	switch b.Compare(v, node.Value) {
	case -1:
		return b.contains(node.left, v)
	case 1:
		return b.contains(node.right, v)
	}
	return true
}

// Contains2 查询是否包含指定元素非递归
func (b *BST) Contains2(v interface{}) bool {
	node := b.root
	for node != nil {
		switch b.Compare(node.Value, v) {
		case 0:
			return true
		case -1:
			node = node.right
		case 1:
			node = node.left
		}
	}
	return false
}

// PreOrder 前序遍历
func (b *BST) PreOrder(f func(v interface{})) {
	b.preOrder(b.root, f)
}

func (b *BST) preOrder(node *Node, f func(v interface{})) {
	if node == nil {
		return
	}

	f(node.Value)
	b.preOrder(node.left, f)
	b.preOrder(node.right, f)
}

// InOrder 中序遍历
func (b *BST) InOrder(f func(v interface{})) {
	b.inOrder(b.root, f)
}

func (b *BST) inOrder(node *Node, f func(v interface{})) {
	if node == nil {
		return
	}

	b.inOrder(node.left, f)
	f(node.Value)
	b.inOrder(node.right, f)
}

// PostOrder 后序遍历
func (b *BST) PostOrder(f func(v interface{})) {
	b.postOrder(b.root, f)
}

func (b *BST) postOrder(node *Node, f func(v interface{})) {
	if node == nil {
		return
	}

	b.postOrder(node.left, f)
	b.postOrder(node.right, f)
	f(node.Value)
}

// PreOrderTraverse 前序遍历非递归
func (b *BST) PreOrderTraverse(f func(v interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.Value)
		return
	}

	s := stack.NewArrayStack()
	s.Push(b.root)
	for s.Len() != 0 {
		node := s.Pop().(*Node)
		f(node.Value)
		if node.right != nil {
			s.Push(node.right)
		}
		if node.left != nil {
			s.Push(node.left)
		}
	}
}

// InOrderTraverse 中序遍历非递归
func (b *BST) InOrderTraverse(f func(v interface{})) {
	if b.root == nil {
		return
	}
	if b.size == 1 {
		f(b.root.Value)
		return
	}

	var node *Node
	var flag = true
	s := stack.NewArrayStack()
	s.Push(b.root)
	for s.Len() != 0 {
		if flag {
			flag = false
			for node = s.Peek().(*Node); node.left != nil; node = node.left {
				s.Push(node.left)
			}
		}
		node = s.Pop().(*Node)
		f(node.Value)
		if node.right != nil {
			s.Push(node.right)
			flag = true
		}
	}
}

// PostOrderTraverse 后序遍历非递归
func (b *BST) PostOrderTraverse(f func(v interface{})) {

}

// LevelOrder 层序遍历
func (b *BST) LevelOrder(f func(v interface{})) {
	if b.root == nil {
		return
	}

	q := queue.NewArrayQueue()
	q.Enqueue(b.root)
	var node *Node
	for q.Len() > 0 {
		node = q.Dequeue().(*Node)
		f(node.Value)
		if node.left != nil {
			q.Enqueue(node.left)
		}
		if node.right != nil {
			q.Enqueue(node.right)
		}
	}
}

// RemoveMax 删除最大元素
func (b *BST) RemoveMax() interface{} {
	switch b.size {
	case 0:
		return nil
	case 1:
		b.size--
		r := b.root
		b.root = nil
		return r.Value
	case 2:
		b.size--
		if b.root.left == nil {
			r := b.root.right
			b.root.right = nil
			return r.Value
		}
		r := b.root
		b.root = b.root.left
		r.left = nil
		return r.Value
	default:
		b.size--
		p := b.root
		n := b.root.right
		for n.right != nil {
			p = n
			n = n.right
		}
		p.right = n.left
		n.left = nil
		return n.Value
	}
}

// RemoveMin 删除最小元素
func (b *BST) RemoveMin() interface{} {
	switch b.size {
	case 0:
		return nil
	case 1:
		b.size--
		r := b.root
		b.root = nil
		return r.Value
	case 2:
		b.size--
		if b.root.left == nil {
			r := b.root
			b.root = b.root.right
			r.right = nil
			return r.Value
		}
		n := b.root.left
		b.root.left = nil
		return n.Value
	default:
		b.size--
		p := b.root
		n := b.root.left
		for n.left != nil {
			p = n
			n = n.left
		}
		p.left = n.right
		n.right = nil
		return n.Value
	}
}

// 删除任意元素
func (b *BST) Remove(e interface{}) bool {
	sz := b.size
	b.root = b.remove(b.root, e)
	return sz > b.size
}

func (b *BST) remove(node *Node, e interface{}) *Node {
	if node == nil {
		return nil
	}

	if n := b.f(e, node.e); n < 0 {
		node.left = b.remove(node.left, e)
	} else if n > 0 {
		node.right = b.remove(node.right, e)
	} else {
		// 删除当前节点的情况
		b.size--
		if node.left == nil { // 返回右节点
			retNode := node.right
			node.right = nil
			return retNode
		}
		if node.right == nil { // 返回左节点
			retNode := node.left
			node.left = nil
			return retNode
		}
		// 待删除节点左右子数均不为空的情况
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		retNode := b.findMin(node.right)
		retNode.right = b.removeMin(node.right)
		b.size++

		retNode.left = node.left

		node.left = nil
		node.right = nil

		return retNode
	}
	return node
}
