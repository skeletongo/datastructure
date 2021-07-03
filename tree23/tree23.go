// 2-3树
package tree23

import (
	"github.com/skeletongo/datastructure/common"
)

type Tree23 struct {
	root    *node
	Compare func(a, b interface{}) int
}

// New 创建左倾红黑树(2-3树)
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func New(compare func(a, b interface{}) int) *Tree23 {
	return &Tree23{Compare: compare}
}

func (t *Tree23) GetSize() int {
	if t.root == nil {
		return 0
	}
	return t.root.n
}

func (t *Tree23) IsEmpty() bool {
	return t.root == nil
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
func (t *Tree23) isBST() bool {
	arr := new([]interface{})
	inOrder(t.root, arr)
	for i := 1; i < len(*arr); i++ {
		if t.Compare((*arr)[i-1], (*arr)[i]) > 0 {
			return false
		}
	}
	return true
}

// isBalanced 判断是不是平衡二叉树(黑平衡)
func (t *Tree23) isBalanced() bool {
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

// balance 维护红节点的位置
// 1.将不正确的3节点和4节点变幻成正确的3节点和4节点
// 2.将4节点分解成两个2节点，将剩余的节点向它的父节点融合
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

func (t *Tree23) put(n *node, key, value interface{}) *node {
	if n == nil {
		return newNode(key, value)
	}

	res := t.Compare(n.key, key)
	if res > 0 {
		n.left = t.put(n.left, key, value)
	} else if res < 0 {
		n.right = t.put(n.right, key, value)
	} else {
		n.value = value
		return n
	}

	return balance(n)
}

func (t *Tree23) Put(key, value interface{}) {
	t.root = t.put(t.root, key, value)
	t.root.color = Black
}

func (t *Tree23) contains(n *node, key interface{}) bool {
	if n == nil {
		return false
	}

	res := t.Compare(n.key, key)
	if res > 0 {
		return t.contains(n.left, key)
	}
	if res < 0 {
		return t.contains(n.right, key)
	}
	return true
}

func (t *Tree23) Contains(key interface{}) bool {
	return t.contains(t.root, key)
}

func (t *Tree23) get(n *node, key interface{}) interface{} {
	if n == nil {
		return nil
	}

	res := t.Compare(n.key, key)
	if res > 0 {
		return t.get(n.left, key)
	}
	if res < 0 {
		return t.get(n.right, key)
	}
	return n.value
}

func (t *Tree23) Get(key interface{}) interface{} {
	return t.get(t.root, key)
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

func (t *Tree23) removeMin(n *node) *node {
	if n.left == nil {
		return nil
	}

	if !isRed(n.left) && !isRed(n.left.left) {
		n = moveRedLeft(n)
	}
	n.left = t.removeMin(n.left)
	return balance(n)
}

func (t *Tree23) RemoveMin() {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		t.root = nil
		return
	}

	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = Red
	}
	t.root = t.removeMin(t.root)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func moveRedRight(n *node) *node {
	colorsFlip(n)
	if !isRed(n.left.left) {
		n = rightRotate(n)
	}
	return n
}

func (t *Tree23) removeMax(n *node) *node {
	if isRed(n.left) {
		n = rightRotate(n)
	}
	if n.right == nil {
		return nil
	}
	if !isRed(n.right) && !isRed(n.right.left) {
		n = moveRedRight(n)
	}
	n.right = t.removeMax(n.right)
	return balance(n)
}

func (t *Tree23) RemoveMax() {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		t.root = nil
		return
	}

	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = Red
	}
	t.root = t.removeMax(t.root)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func minKey(n *node) interface{} {
	cur := n
	for cur.left != nil {
		cur = cur.left
	}
	return cur.key
}

func (t *Tree23) remove(n *node, key interface{}) *node {
	if n == nil {
		return nil
	}

	if t.Compare(n.key, key) > 0 {
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
		}
		n.left = t.remove(n.left, key)
	} else {
		if isRed(n.left) {
			n = rightRotate(n)
		}
		if t.Compare(n.key, key) == 0 && n.right == nil {
			return nil
		}
		if !isRed(n.right) && !isRed(n.right.left) {
			n = moveRedRight(n)
		}
		if t.Compare(n.key, key) == 0 {
			n.key = minKey(n.right)
			n.value = t.get(n.right, n.key)
			n.right = t.removeMin(n.right)
		} else {
			n.right = t.remove(n.right, key)
		}
	}
	return balance(n)
}

func (t *Tree23) Remove(key interface{}) {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		if t.Compare(t.root.key, key) == 0 {
			t.root = nil
		}
		return
	}

	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = Red
	}
	t.root = t.remove(t.root, key)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func (t *Tree23) Range(f func(key, value interface{})) {
	common.PreOrder(t.root, f)
}

func (t *Tree23) String() string {
	return common.PrePrint(t.root)
}
