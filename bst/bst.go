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

func (b *BST) add(node *Node, v interface{}) *Node {
	if node == nil {
		b.size++
		return &Node{Value: v}
	}

	r := b.Compare(v, node.Value)
	if r < 0 {
		node.left = b.add(node.left, v)
	} else if r > 0 {
		node.right = b.add(node.right, v)
	}
	return node
}

// Add 添加新节点
func (b *BST) Add(v interface{}) {
	b.root = b.add(b.root, v)
}

// AddNR 添加新节点非递归
func (b *BST) AddNR(v interface{}) {
	if b.size == 0 {
		b.size++
		b.root = &Node{Value: v}
		return
	}

	node := b.root
	for {
		r := b.Compare(v, node.Value)
		switch {
		case r < 0:
			if node.left == nil {
				b.size++
				node.left = &Node{Value: v}
				return
			}
			node = node.left
		case r > 0:
			if node.right == nil {
				b.size++
				node.right = &Node{Value: v}
				return
			}
			node = node.right
		default:
			return
		}
	}
}

func (b *BST) contains(node *Node, v interface{}) bool {
	if node == nil {
		return false
	}

	r := b.Compare(v, node.Value)
	if r < 0 {
		return b.contains(node.left, v)
	}
	if r > 0 {
		return b.contains(node.right, v)
	}
	return true
}

// Contains 查询是否包含指定元素
func (b *BST) Contains(v interface{}) bool {
	return b.contains(b.root, v)
}

// ContainsNR 查询是否包含指定元素非递归
func (b *BST) ContainsNR(v interface{}) bool {
	node := b.root
	for node != nil {
		r := b.Compare(v, node.Value)
		if r == 0 {
			return true
		}
		if r < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return false
}

func (b *BST) preOrder(node *Node, f func(v interface{})) {
	if node == nil {
		return
	}

	f(node.Value)
	b.preOrder(node.left, f)
	b.preOrder(node.right, f)
}

// PreOrder 前序遍历
func (b *BST) PreOrder(f func(v interface{})) {
	b.preOrder(b.root, f)
}

func (b *BST) inOrder(node *Node, f func(v interface{})) {
	if node == nil {
		return
	}

	b.inOrder(node.left, f)
	f(node.Value)
	b.inOrder(node.right, f)
}

// InOrder 中序遍历
func (b *BST) InOrder(f func(v interface{})) {
	b.inOrder(b.root, f)
}

func (b *BST) postOrder(node *Node, f func(v interface{})) {
	if node == nil {
		return
	}

	b.postOrder(node.left, f)
	b.postOrder(node.right, f)
	f(node.Value)
}

// PostOrder 后序遍历
func (b *BST) PostOrder(f func(v interface{})) {
	b.postOrder(b.root, f)
}

// PreOrderNR 前序遍历非递归
func (b *BST) PreOrderNR(f func(v interface{})) {
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

// InOrderNR 中序遍历非递归
func (b *BST) InOrderNR(f func(v interface{})) {
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

// PostOrderNR 后序遍历非递归
func (b *BST) PostOrderNR(f func(v interface{})) {

}

// PreOrderNRC 前序遍历非递归经典版
func (b *BST) PreOrderNRC(f func(v interface{})) {

}

// InOrderNRC 中序遍历非递归经典版
func (b *BST) InOrderNRC(f func(v interface{})) {

}

// PostOrderNRC 后序遍历非递归经典版
func (b *BST) PostOrderNRC(f func(v interface{})) {

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

// findMin 寻找当前节点中的最小值节点
func findMin(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return findMin(node.left)
}

// removeMin 删除当前节点中的最小值节点并返回当前节点删除最小值节点后的根节点
func (b *BST) removeMin(node *Node) *Node {
	if node.left == nil {
		b.size--
		ret := node.right
		node.right = nil
		return ret
	}
	node.left = b.removeMin(node.left)
	return node
}

// RemoveMin 删除最小值节点并返回删除的最小值
func (b *BST) RemoveMin() interface{} {
	if b.size == 0 {
		panic("no data")
	}
	node := findMin(b.root)
	b.root = b.removeMin(b.root)
	return node.Value
}

// findMax 寻找当前节点中的最大值节点
func findMax(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return findMax(node.right)
}

// removeMax 删除当前节点中的最大值节点并返回当前节点删除最大值节点后的根节点
func (b *BST) removeMax(node *Node) *Node {
	if node.right == nil {
		b.size--
		ret := node.left
		node.left = nil
		return ret
	}
	node.right = b.removeMax(node.right)
	return node
}

// RemoveMax 删除最大值节点并返回删除的最大值
func (b *BST) RemoveMax() interface{} {
	if b.size == 0 {
		panic("no data")
	}

	node := findMax(b.root)
	b.root = b.removeMax(b.root)
	return node.Value
}

// RemoveMinNR 删除最小值节点并返回删除的最小值
func (b *BST) RemoveMinNR() interface{} {
	switch b.size {
	case 0:
		panic("no data")
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
		s := p.left
		for s.left != nil {
			p = s
			s = s.left
		}
		p.left = s.right
		s.right = nil
		return s.Value
	}
}

// RemoveMaxNR 删除最大值节点并返回删除的最大值
func (b *BST) RemoveMaxNR() interface{} {
	switch b.size {
	case 0:
		panic("no data")
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

// remove 从当前节点中删除一个特定的值节点并返回删除节点后的当前节点的根节点
func (b *BST) remove(node *Node, v interface{}) *Node {
	if node == nil {
		return nil
	}

	r := b.Compare(v, node.Value)
	if r < 0 {
		node.left = b.remove(node.left, v)
		return node
	}
	if r > 0 {
		node.right = b.remove(node.right, v)
		return node
	}

	if node.left == nil {
		b.size--
		ret := node.right
		node.right = nil
		return ret
	}
	if node.right == nil {
		b.size--
		ret := node.left
		node.left = nil
		return ret
	}

	n := findMin(node.right)
	n.right = b.removeMin(node.right)
	n.left = node.left
	node.left = nil
	node.right = nil
	return n
}

// Remove 删除值节点
func (b *BST) Remove(v interface{}) {
	b.root = b.remove(b.root, v)
}
