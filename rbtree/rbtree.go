package rbtree

import (
	"dataStructure/common"
	"fmt"
)

const (
	Red   = true
	Black = false
)

type node struct {
	left, right *node
	color       bool
	n           int         // 以当前节点为根的树的节点数量
	key, value  interface{} // 存储映射的键值对
}

func newNode(key, value interface{}) *node {
	return &node{
		key:   key,
		value: value,
		color: Red,
		n:     1,
	}
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

type RBTree struct {
	root    *node
	Compare func(a, b interface{}) int
}

// New 创建左倾红黑树(2-3树)
// 参数 f 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func New(f func(a, b interface{}) int) *RBTree {
	return &RBTree{Compare: f}
}

func (r *RBTree) GetSize() int {
	if r.root == nil {
		return 0
	}
	return r.root.n
}

func (r *RBTree) IsEmpty() bool {
	return r.root == nil
}

func inOrder(n *node, list *[]interface{}) {
	if n == nil {
		return
	}
	inOrder(n.left, list)
	*list = append(*list, n.key)
	inOrder(n.right, list)
}

// isBST 判断是不是二分搜索树
func (r *RBTree) isBST() bool {
	list := new([]interface{})
	inOrder(r.root, list)
	for i := 1; i < len(*list); i++ {
		if r.Compare((*list)[i-1], (*list)[i]) > 0 {
			return false
		}
	}
	return true
}

// isBalanced 判断是不是平衡二叉树(黑平衡)
func (r *RBTree) isBalanced() bool {
	return true
}

func size(n *node) int {
	if n == nil {
		return 0
	}
	return n.n
}

// 隐含条件，n不为nil
func getSize(n *node) int {
	return size(n.left) + size(n.right) + 1
}

// isRed 是否为红节点
func isRed(n *node) bool {
	if n == nil {
		return Black
	}
	return n.color // n.color == red 不用等号判断需要定义Red为true
}

// leftRotate 左旋转
//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func leftRotate(n *node) *node {
	x := n.right
	n.right = x.left
	x.left = n
	x.color = n.color
	n.color = Red
	x.n = n.n
	n.n = getSize(n)
	return x
}

// rightRotate 右旋转
//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func rightRotate(n *node) *node {
	x := n.left
	n.left = x.right
	x.right = n
	x.color = n.color
	n.color = Red
	x.n = n.n
	n.n = getSize(n)
	return x
}

// flipColors 颜色翻转
func flipColors(n *node) {
	n.color = Red
	n.left.color = Black
	n.right.color = Black
}

// balance 维护红节点的位置，也是分解4节点(颜色翻转)的过程
// 隐含条件，n不为nil
func balance(n *node) *node {
	if isRed(n.right) && !isRed(n.left) {
		n = leftRotate(n)
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = rightRotate(n)
	}
	if isRed(n.left) && isRed(n.right) {
		flipColors(n)
	}
	n.n = getSize(n)
	return n
}

func (r *RBTree) put(n *node, key, value interface{}) *node {
	if n == nil {
		return newNode(key, value)
	}

	res := r.Compare(n.key, key)
	if res > 0 {
		n.left = r.put(n.left, key, value)
	} else if res < 0 {
		n.right = r.put(n.right, key, value)
	} else {
		n.value = value
		return n
	}

	return balance(n)
}

func (r *RBTree) Put(key, value interface{}) {
	r.root = r.put(r.root, key, value)
	r.root.color = Black
}

func (r *RBTree) contains(n *node, key interface{}) bool {
	if n == nil {
		return false
	}

	res := r.Compare(n.key, key)
	if res > 0 {
		return r.contains(n.left, key)
	}
	if res < 0 {
		return r.contains(n.right, key)
	}
	return true
}

func (r *RBTree) Contains(key interface{}) bool {
	return r.contains(r.root, key)
}

func (r *RBTree) get(n *node, key interface{}) interface{} {
	if n == nil {
		return nil
	}

	res := r.Compare(n.key, key)
	if res > 0 {
		return r.get(n.left, key)
	}
	if res < 0 {
		return r.get(n.right, key)
	}
	return n.value
}

func (r *RBTree) Get(key interface{}) interface{} {
	return r.get(r.root, key)
}

// colorsFlip 颜色翻转
func colorsFlip(n *node) {
	n.color = Black
	n.left.color = Red
	n.right.color = Red
}

func moveRedLeft(n *node) *node {
	colorsFlip(n)
	if isRed(n.right.left) {
		n.right = rightRotate(n.right)
		n = leftRotate(n)
	}
	return n
}

func (r *RBTree) removeMin(n *node) *node {
	if n.left == nil {
		return nil
	}

	if !isRed(n.left) && !isRed(n.left.left) {
		n = moveRedLeft(n)
	}
	n.left = r.removeMin(n.left)
	return balance(n)
}

func (r *RBTree) RemoveMin() {
	if r.root == nil {
		return
	}

	if r.root.n == 1 {
		r.root = nil
		return
	}

	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = Red
	}
	r.root = r.removeMin(r.root)
	if !r.IsEmpty() {
		r.root.color = Black
	}
}

func moveRedRight(n *node) *node {
	colorsFlip(n)
	if !isRed(n.left.left) {
		n = rightRotate(n)
	}
	return n
}

func (r *RBTree) removeMax(n *node) *node {
	if isRed(n.left) {
		n = rightRotate(n)
	}
	if n.right == nil {
		return nil
	}
	if !isRed(n.right) && !isRed(n.right.left) {
		n = moveRedRight(n)
	}
	n.right = r.removeMax(n.right)
	return balance(n)
}

func (r *RBTree) RemoveMax() {
	if r.root == nil {
		return
	}

	if r.root.n == 1 {
		r.root = nil
		return
	}

	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = Red
	}
	r.root = r.removeMax(r.root)
	if !r.IsEmpty() {
		r.root.color = Black
	}
}

func minKey(n *node) interface{} {
	cur := n
	for cur.left != nil {
		cur = cur.left
	}
	return cur.key
}

func (r *RBTree) remove(n *node, key interface{}) *node {
	if n == nil {
		return nil
	}

	res := r.Compare(n.key, key)
	if res > 0 {
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
		}
		n.left = r.remove(n.left, key)
	} else {
		if isRed(n.left) {
			n = rightRotate(n)
		}
		if res == 0 && n.right == nil {
			return nil
		}
		if !isRed(n.right) && !isRed(n.right.left) {
			n = moveRedRight(n)
		}
		if res == 0 {
			n.key = minKey(n.right)
			n.value = r.get(n.right, n.key)
			n.right = r.removeMin(n.right)
		} else {
			n.right = r.remove(n.right, key)
		}
	}
	return balance(n)
}

func (r *RBTree) Remove(key interface{}) {
	if r.root == nil {
		return
	}

	if r.root.n == 1 {
		if r.Compare(r.root.key, key) == 0 {
			r.root = nil
		}
		return
	}

	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = Red
	}
	r.root = r.remove(r.root, key)
	if !r.IsEmpty() {
		r.root.color = Black
	}
}

func (r *RBTree) String() string {
	return common.PrePrint(r.root)
}
