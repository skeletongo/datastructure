// 二分搜索树的查询,添加,删除操作的时间复杂度为O(h) 最差时间复杂度O(n)链表 最佳时间复杂O(logN)满二叉树
// 满二叉树：除了最大层节点，其他节点都有两个子节点
// 完全二叉树：按每层从左到右的位置添加新节点
// 平衡二叉树：所有叶子节点所在的层数的差值的绝对值不能大于1
package bst

import (
	"fmt"

	"github.com/skeletongo/dataStructure/common"
	"github.com/skeletongo/dataStructure/queue"
	"github.com/skeletongo/dataStructure/stack"
)

type node struct {
	left, right *node
	key, value  interface{}
}

func newNode(key, value interface{}) *node {
	return &node{key: key, value: value}
}

func (n *node) GetLeftNode() common.INode {
	return n.left
}

func (n *node) GetRightNode() common.INode {
	return n.right
}

func (n *node) GetValue() interface{} {
	return fmt.Sprintf("%v: %v", n.key, n.value)
}

// BST 二分搜索树
type BST struct {
	// 根节点
	root *node
	// 节点数量
	size int
	// 元素比较方法
	Compare func(a, b interface{}) int
}

// New 创建二分搜索树
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func New(compare func(a, b interface{}) int) *BST {
	return &BST{Compare: compare}
}

// GetSize 获取节点总数
func (b *BST) GetSize() int {
	return b.size
}

// IsEmpty 是否为空树
func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) add(n *node, key, value interface{}) *node {
	if n == nil {
		b.size++
		return newNode(key, value)
	}

	res := b.Compare(key, n.key)
	if res < 0 {
		n.left = b.add(n.left, key, value)
	} else if res > 0 {
		n.right = b.add(n.right, key, value)
	}
	return n
}

// Add 添加新节点
func (b *BST) Add(key, value interface{}) {
	b.root = b.add(b.root, key, value)
}

// AddNR 添加新节点非递归
func (b *BST) AddNR(key, value interface{}) {
	if b.size == 0 {
		b.size++
		b.root = newNode(key, value)
		return
	}

	n := b.root
	for {
		r := b.Compare(key, n.key)
		switch {
		case r < 0:
			if n.left == nil {
				b.size++
				n.left = newNode(key, value)
				return
			}
			n = n.left
		case r > 0:
			if n.right == nil {
				b.size++
				n.right = newNode(key, value)
				return
			}
			n = n.right
		default:
			return
		}
	}
}

func (b *BST) contains(n *node, key interface{}) bool {
	if n == nil {
		return false
	}

	r := b.Compare(key, n.key)
	if r < 0 {
		return b.contains(n.left, key)
	}
	if r > 0 {
		return b.contains(n.right, key)
	}
	return true
}

// Contains 查询是否包含指定元素
func (b *BST) Contains(key interface{}) bool {
	return b.contains(b.root, key)
}

// ContainsNR 查询是否包含指定元素非递归
func (b *BST) ContainsNR(key interface{}) bool {
	n := b.root
	for n != nil {
		r := b.Compare(key, n.key)
		if r == 0 {
			return true
		}
		if r < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}
	return false
}

// findMin 寻找当前节点中的最小值节点
func findMin(n *node) *node {
	if n.left == nil {
		return n
	}
	return findMin(n.left)
}

// removeMin 删除当前节点中的最小值节点并返回当前节点删除最小值节点后的根节点
func (b *BST) removeMin(n *node) *node {
	if n.left == nil {
		b.size--
		ret := n.right
		n.right = nil
		return ret
	}
	n.left = b.removeMin(n.left)
	return n
}

// RemoveMin 删除最小值节点并返回删除的最小值
func (b *BST) RemoveMin() interface{} {
	if b.size == 0 {
		panic("no data")
	}
	n := findMin(b.root)
	b.root = b.removeMin(b.root)
	return n.key
}

// findMax 寻找当前节点中的最大值节点
func findMax(n *node) *node {
	if n.right == nil {
		return n
	}
	return findMax(n.right)
}

// removeMax 删除当前节点中的最大值节点并返回当前节点删除最大值节点后的根节点
func (b *BST) removeMax(n *node) *node {
	if n.right == nil {
		b.size--
		ret := n.left
		n.left = nil
		return ret
	}
	n.right = b.removeMax(n.right)
	return n
}

// RemoveMax 删除最大值节点并返回删除的最大值
func (b *BST) RemoveMax() interface{} {
	if b.size == 0 {
		panic("no data")
	}

	n := findMax(b.root)
	b.root = b.removeMax(b.root)
	return n.key
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
		return r.key
	case 2:
		b.size--
		if b.root.left == nil {
			r := b.root
			b.root = b.root.right
			r.right = nil
			return r.key
		}
		n := b.root.left
		b.root.left = nil
		return n.key
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
		return s.key
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
		return r.key
	case 2:
		b.size--
		if b.root.left == nil {
			r := b.root.right
			b.root.right = nil
			return r.value
		}
		n := b.root
		b.root = b.root.left
		n.left = nil
		return n.key
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
		return n.key
	}
}

// remove 从当前节点中删除一个特定的值节点并返回删除节点后的当前节点的根节点
func (b *BST) remove(n *node, key interface{}) *node {
	if n == nil {
		return nil
	}

	r := b.Compare(key, n.key)
	if r < 0 {
		n.left = b.remove(n.left, key)
		return n
	}
	if r > 0 {
		n.right = b.remove(n.right, key)
		return n
	}

	if n.left == nil {
		b.size--
		ret := n.right
		n.right = nil
		return ret
	}
	if n.right == nil {
		b.size--
		ret := n.left
		n.left = nil
		return ret
	}

	min := findMin(n.right)
	min.right = b.removeMin(n.right)
	min.left = n.left
	n.left = nil
	n.right = nil
	return min
}

// Remove 删除值节点
func (b *BST) Remove(key interface{}) {
	b.root = b.remove(b.root, key)
}

func (b *BST) set(n *node, key, value interface{}) *node {
	if n == nil {
		b.size++
		return newNode(key, value)
	}

	res := b.Compare(n.key, key)
	if res < 0 {
		n.right = b.set(n.right, key, value)
		return n
	}
	if res > 0 {
		n.left = b.set(n.left, key, value)
		return n
	}
	n.value = value
	return n
}

// Set 添加元素或修改元素对应的值
func (b *BST) Set(key, value interface{}) {
	b.root = b.set(b.root, key, value)
}

func (b *BST) get(n *node, key interface{}) *node {
	if n == nil {
		return nil
	}

	res := b.Compare(n.key, key)
	if res < 0 {
		return b.get(n.right, key)
	}
	if res > 0 {
		return b.get(n.left, key)
	}
	return n
}

// Get 获取元素值
func (b *BST) Get(key interface{}) interface{} {
	n := b.get(b.root, key)
	if n == nil {
		return nil
	}
	return n.value
}

func preOrder(n *node, f func(key, value interface{})) {
	if n == nil {
		return
	}

	f(n.key, n.value)
	preOrder(n.left, f)
	preOrder(n.right, f)
}

// PreOrder 前序遍历
func (b *BST) PreOrder(f func(key, value interface{})) {
	preOrder(b.root, f)
}

func inOrder(n *node, f func(key, value interface{})) {
	if n == nil {
		return
	}

	inOrder(n.left, f)
	f(n.key, n.value)
	inOrder(n.right, f)
}

// InOrder 中序遍历
func (b *BST) InOrder(f func(key, value interface{})) {
	inOrder(b.root, f)
}

func postOrder(n *node, f func(key, value interface{})) {
	if n == nil {
		return
	}

	postOrder(n.left, f)
	postOrder(n.right, f)
	f(n.key, n.value)
}

// PostOrder 后序遍历
func (b *BST) PostOrder(f func(key, value interface{})) {
	postOrder(b.root, f)
}

// PreOrderNR 前序遍历非递归
func (b *BST) PreOrderNR(f func(key, value interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.key, b.root.value)
		return
	}

	s := stack.NewArrayStack()
	s.Push(b.root)
	for s.Len() != 0 {
		n := s.Pop().(*node)
		f(n.key, n.value)
		if n.right != nil {
			s.Push(n.right)
		}
		if n.left != nil {
			s.Push(n.left)
		}
	}
}

// InOrderNR 中序遍历非递归
func (b *BST) InOrderNR(f func(key, value interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.key, b.root.value)
		return
	}

	var n *node
	var flag = true
	s := stack.NewArrayStack()
	s.Push(b.root)
	for s.Len() != 0 {
		if flag {
			flag = false
			for n = s.Peek().(*node); n.left != nil; n = n.left {
				s.Push(n.left)
			}
		}
		n = s.Pop().(*node)
		f(n.key, n.value)
		if n.right != nil {
			s.Push(n.right)
			flag = true
		}
	}
}

// PostOrderNR 后序遍历非递归(双栈方式)
func (b *BST) PostOrderNR(f func(key, value interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.key, b.root.value)
		return
	}

	var n *node
	s1 := stack.NewArrayStack()
	s2 := stack.NewArrayStack()
	s1.Push(b.root)
	for s1.Len() > 0 {
		n = s1.Pop().(*node)
		s2.Push(n)
		if n.left != nil {
			s1.Push(n.left)
		}
		if n.right != nil {
			s1.Push(n.right)
		}
	}
	for s2.Len() > 0 {
		n = s2.Pop().(*node)
		f(n.key, n.value)
	}
}

// PreOrderNRC 前序遍历非递归经典版
func (b *BST) PreOrderNRC(f func(key, value interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.key, b.root.value)
		return
	}

	s := stack.NewArrayStack()
	n := b.root
	for s.Len() > 0 || n != nil {
		if n != nil {
			f(n.key, n.value)
			s.Push(n)
			n = n.left
		} else {
			n = s.Pop().(*node).right
		}
	}
}

// InOrderNRC 中序遍历非递归经典版
func (b *BST) InOrderNRC(f func(key, value interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.value, b.root.value)
		return
	}

	s := stack.NewArrayStack()
	n := b.root
	for s.Len() > 0 || n != nil {
		if n != nil {
			s.Push(n)
			n = n.left
		} else {
			n = s.Pop().(*node)
			f(n.key, n.value)
			n = n.right
		}
	}
}

// PostOrderNRC 后序遍历非递归经典版
func (b *BST) PostOrderNRC(f func(key, value interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.key, b.root.value)
		return
	}

	var ok bool
	s := stack.NewArrayStack()
	m := map[*node]struct{}{}
	n := b.root
	for s.Len() > 0 || n != nil {
		if n != nil {
			s.Push(n)
			n = n.left
		} else {
			n = s.Peek().(*node)
			if _, ok = m[n]; ok {
				s.Pop()
				f(n.key, n.value)
				n = nil
			} else {
				m[n] = struct{}{}
				n = n.right
			}
		}
	}
}

// LevelOrder 层序遍历
func (b *BST) LevelOrder(f func(key, value interface{})) {
	if b.root == nil {
		return
	}

	q := queue.NewArrayQueue()
	q.Enqueue(b.root)
	var n *node
	for q.Len() > 0 {
		n = q.Dequeue().(*node)
		f(n.key, n.value)
		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
}

func (b *BST) String() string {
	return common.PrePrint(b.root)
}
