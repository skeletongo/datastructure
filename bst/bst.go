// 二分搜索树的查询,添加,删除操作的时间复杂度为O(h) 最差时间复杂度O(n)链表 最佳时间复杂O(logN)满二叉树
// 满二叉树：除了最大层节点，其他节点都有两个子节点
// 完全二叉树：按每层从左到右的位置添加新节点
// 平衡二叉树：所有叶子节点所在的层数的差值的绝对值不能大于1
package bst

import (
	"github.com/skeletongo/datastructure/common"
	"github.com/skeletongo/datastructure/stack"
)

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

func (b *BST) add(n *node, value interface{}) *node {
	if n == nil {
		b.size++
		return newNode(value)
	}

	res := b.Compare(value, n.value)
	if res < 0 {
		n.left = b.add(n.left, value)
	} else if res > 0 {
		n.right = b.add(n.right, value)
	}
	return n
}

// Add 添加新节点
func (b *BST) Add(value interface{}) {
	b.root = b.add(b.root, value)
}

// AddNR 添加新节点非递归
func (b *BST) AddNR(value interface{}) {
	if b.size == 0 {
		b.size++
		b.root = newNode(value)
		return
	}

	n := b.root
	for {
		r := b.Compare(value, n.value)
		switch {
		case r < 0:
			if n.left == nil {
				b.size++
				n.left = newNode(value)
				return
			}
			n = n.left
		case r > 0:
			if n.right == nil {
				b.size++
				n.right = newNode(value)
				return
			}
			n = n.right
		default:
			return
		}
	}
}

func (b *BST) contains(n *node, value interface{}) bool {
	if n == nil {
		return false
	}

	r := b.Compare(value, n.value)
	if r < 0 {
		return b.contains(n.left, value)
	}
	if r > 0 {
		return b.contains(n.right, value)
	}
	return true
}

// Contains 查询是否包含指定元素
func (b *BST) Contains(value interface{}) bool {
	return b.contains(b.root, value)
}

// ContainsNR 查询是否包含指定元素非递归
func (b *BST) ContainsNR(value interface{}) bool {
	n := b.root
	for n != nil {
		r := b.Compare(value, n.value)
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
	return n.value
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
	return n.value
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
		return r.value
	case 2:
		b.size--
		if b.root.left == nil {
			r := b.root
			b.root = b.root.right
			r.right = nil
			return r.value
		}
		n := b.root.left
		b.root.left = nil
		return n.value
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
		return s.value
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
		return r.value
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
		return n.value
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
		return n.value
	}
}

// remove 从当前节点中删除一个特定的值节点并返回删除节点后的当前节点的根节点
func (b *BST) remove(n *node, value interface{}) *node {
	if n == nil {
		return nil
	}

	r := b.Compare(value, n.value)
	if r < 0 {
		n.left = b.remove(n.left, value)
		return n
	}
	if r > 0 {
		n.right = b.remove(n.right, value)
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

func (b *BST) set(n *node, value interface{}) *node {
	if n == nil {
		b.size++
		return newNode(value)
	}

	res := b.Compare(n.value, value)
	if res < 0 {
		n.right = b.set(n.right, value)
		return n
	}
	if res > 0 {
		n.left = b.set(n.left, value)
		return n
	}
	n.value = value
	return n
}

// Set 添加元素或修改元素对应的值
func (b *BST) Set(value interface{}) {
	b.root = b.set(b.root, value)
}

func (b *BST) get(n *node, value interface{}) *node {
	if n == nil {
		return nil
	}

	res := b.Compare(n.value, value)
	if res < 0 {
		return b.get(n.right, value)
	}
	if res > 0 {
		return b.get(n.left, value)
	}
	return n
}

// Get 获取元素值
func (b *BST) Get(value interface{}) interface{} {
	n := b.get(b.root, value)
	if n == nil {
		return nil
	}
	return n.value
}

// PreOrder 前序遍历
func (b *BST) PreOrder(f func(v interface{})) {
	common.PreOrder(b.root, f)
}

// InOrder 中序遍历
func (b *BST) InOrder(f func(v interface{})) {
	common.InOrder(b.root, f)
}

// PostOrder 后序遍历
func (b *BST) PostOrder(f func(v interface{})) {
	common.PostOrder(b.root, f)
}

// PreOrderNR 前序遍历非递归
func (b *BST) PreOrderNR(f func(value interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.value)
		return
	}

	s := stack.NewArrayStack()
	s.Push(b.root)
	for s.Len() != 0 {
		n := s.Pop().(*node)
		f(n.value)
		if n.right != nil {
			s.Push(n.right)
		}
		if n.left != nil {
			s.Push(n.left)
		}
	}
}

// InOrderNR 中序遍历非递归
func (b *BST) InOrderNR(f func(interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.value)
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
		f(n.value)
		if n.right != nil {
			s.Push(n.right)
			flag = true
		}
	}
}

// PostOrderNR 后序遍历非递归(双栈方式)
func (b *BST) PostOrderNR(f func(interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.value)
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
		f(n.value)
	}
}

// PreOrderNRC 前序遍历非递归经典版
func (b *BST) PreOrderNRC(f func(interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.value)
		return
	}

	s := stack.NewArrayStack()
	n := b.root
	for s.Len() > 0 || n != nil {
		if n != nil {
			f(n.value)
			s.Push(n)
			n = n.left
		} else {
			n = s.Pop().(*node).right
		}
	}
}

// InOrderNRC 中序遍历非递归经典版
func (b *BST) InOrderNRC(f func(interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.value)
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
			f(n.value)
			n = n.right
		}
	}
}

// PostOrderNRC 后序遍历非递归经典版
func (b *BST) PostOrderNRC(f func(interface{})) {
	if b.size == 0 {
		return
	}
	if b.size == 1 {
		f(b.root.value)
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
				f(n.value)
				n = nil
			} else {
				m[n] = struct{}{}
				n = n.right
			}
		}
	}
}

// LevelOrder 层序遍历
func (b *BST) LevelOrder(f func(interface{})) {
	common.LevelOrder(b.root, f)
}

func (b *BST) Range(f func(interface{})) {
	common.PreOrder(b.root, f)
}

// Img 生成图片
func (b *BST) Img(filename string) error {
	if filename == "" {
		filename = "bst"
	}
	if b.GetSize() > 0 {
		return common.BSTSvg(b.root, filename)
	}
	return nil
}

func (b *BST) String() string {
	return common.PrePrintBST(b.root)
}
