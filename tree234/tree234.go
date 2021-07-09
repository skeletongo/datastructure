// 红黑树（2-3-4树）
package tree234

import (
	"github.com/skeletongo/datastructure/common"
)

type Tree234 struct {
	root    *node
	Compare func(a, b interface{}) int
}

// New 创建红黑树(2-3-4树)
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func New(compare func(a, b interface{}) int) *Tree234 {
	return &Tree234{Compare: compare}
}

func (t *Tree234) GetSize() int {
	if t.root == nil {
		return 0
	}
	return t.root.n
}

func (t *Tree234) IsEmpty() bool {
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
func (t *Tree234) isBST() bool {
	arr := new([]interface{})
	inOrder(t.root, arr)
	for i := 1; i < len(*arr); i++ {
		if t.Compare((*arr)[i-1], (*arr)[i]) > 0 {
			return false
		}
	}
	return true
}

func preOrder(n *node) bool {
	if n == nil {
		return true
	}

	if isRed(n) && (isRed(n.left) || isRed(n.right)) {
		return false
	}
	if isRed(n) && n.left != nil && n.right == nil {
		return false
	}
	if isRed(n) && n.left == nil && n.right != nil {
		return false
	}
	if preOrder(n.left) {
		return preOrder(n.right)
	}
	return false
}

// isBalanced 判断是不是平衡二叉树(黑平衡)
func (t *Tree234) isBalanced() bool {
	return preOrder(t.root)
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

func (t *Tree234) put(n *node, key, value interface{}) *node {
	if n == nil {
		return newNode(key, value)
	}

	if isRed(n.left) && isRed(n.right) {
		flipColors(n)
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

	// 路径回溯维护平衡性
	if isRed(n.left) {
		if isRed(n.left.right) {
			n.left = leftRotate(n.left)
		}
		if isRed(n.left.left) {
			n = rightRotate(n)
		}
	} else if isRed(n.right) {
		if isRed(n.right.left) {
			n.right = rightRotate(n.right)
		}
		if isRed(n.right.right) {
			n = leftRotate(n)
		}
	}
	n.n = getSize(n)
	return n
}

func (t *Tree234) Put(key, value interface{}) {
	t.root = t.put(t.root, key, value)
	t.root.color = Black
}

func (t *Tree234) contains(n *node, key interface{}) bool {
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

func (t *Tree234) Contains(key interface{}) bool {
	return t.contains(t.root, key)
}

func (t *Tree234) get(n *node, key interface{}) interface{} {
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

func (t *Tree234) Get(key interface{}) interface{} {
	return t.get(t.root, key)
}

// colorsFlip 颜色翻转
func colorsFlip(n *node) {
	n.color = Black
	n.left.color = Red
	n.right.color = Red
}

// n为红节点，并且左右子节点都是黑色，都不为空
func moveRedLeft(n *node) *node {
	colorsFlip(n)
	if isRed(n.right.left) {
		n.right = rightRotate(n.right)
		n = leftRotate(n)
	}
	return n
}

// n为红节点，并且左右子节点都是黑色，都不为空
func moveRedRight(n *node) *node {
	colorsFlip(n)
	if isRed(n.left.right) {
		n.left = leftRotate(n.left)
		n = rightRotate(n)
	}
	return n
}

func (t *Tree234) removeMin(n *node) *node {
	if n.left == nil {
		if n.right != nil {
			n = leftRotate(n)
		} else {
			return nil
		}
	}
	switch {
	case isRed(n.left) || isRed(n.left.left) || isRed(n.left.right):
	case isRed(n.right):
		n = leftRotate(n)
	default:
		n = moveRedLeft(n)
	}
	n.left = t.removeMin(n.left)

	// 路径回溯维护平衡性
	if isRed(n.left) && isRed(n.left.left) {
		if !isRed(n.right) {
			n = rightRotate(n)
		} else {
			flipColors(n)
		}
	} else if isRed(n.right) && isRed(n.right.right) {
		if !isRed(n.left) {
			n = leftRotate(n)
		} else {
			flipColors(n)
		}
	}
	n.n = getSize(n)
	return n
}

func (t *Tree234) RemoveMin() {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		t.root = nil
		return
	}

	if !isRed(t.root.left) && isRed(t.root.right) {
		t.root.color = Red
	}
	t.root = t.removeMin(t.root)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func (t *Tree234) removeMax(n *node) *node {
	if n.right == nil {
		if n.left != nil {
			n = rightRotate(n)
		} else {
			return nil
		}
	}
	switch {
	case isRed(n.right) || isRed(n.right.right) || isRed(n.right.left):
	case isRed(n.left):
		n = rightRotate(n)
	default:
		n = moveRedRight(n)
	}
	n.right = t.removeMax(n.right)

	// 路径回溯维护平衡性
	if isRed(n.left) && isRed(n.left.left) {
		if !isRed(n.right) {
			n = rightRotate(n)
		} else {
			flipColors(n)
		}
	} else if isRed(n.right) && isRed(n.right.right) {
		if !isRed(n.left) {
			n = leftRotate(n)
		} else {
			flipColors(n)
		}
	}
	n.n = getSize(n)
	return n
}

func (t *Tree234) RemoveMax() {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		t.root = nil
		return
	}

	if !isRed(t.root.left) && isRed(t.root.right) {
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

func maxKey(n *node) interface{} {
	cur := n
	for cur.right != nil {
		cur = cur.right
	}
	return cur.key
}

func (t *Tree234) remove(n *node, key interface{}) *node {
	res := t.Compare(n.key, key)
	if res > 0 {
		if n.left == nil {
			n = leftRotate(n)
		}
		switch {
		case isRed(n.left) || isRed(n.left.left) || isRed(n.left.right):
		case isRed(n.right):
			n = leftRotate(n)
		default:
			n = moveRedLeft(n)
		}
		n.left = t.remove(n.left, key)
	} else if res < 0 {
		if n.right == nil {
			n = rightRotate(n)
		}
		switch {
		case isRed(n.right) || isRed(n.right.right) || isRed(n.right.left):
		case isRed(n.left):
			n = rightRotate(n)
		default:
			n = moveRedRight(n)
		}
		n.right = t.remove(n.right, key)
	} else {
		if n.left == nil && n.right == nil {
			return nil
		}
		if !isRed(n.left) && !isRed(n.right) {
			colorsFlip(n)
		}
		if isRed(n.right) {
			n.key = minKey(n.right)
			n.value = t.get(n.right, n.key)
			n.right = t.removeMin(n.right)
		} else {
			n.key = maxKey(n.left)
			n.value = t.get(n.left, n.key)
			n.left = t.removeMax(n.left)
		}
	}

	// 路径回溯维护平衡性
	if isRed(n.left) && !isRed(n.right) {
		if isRed(n.left.right) && !isRed(n.left.left) {
			n.left = leftRotate(n.left)
		}
		if isRed(n.left.left) {
			n = rightRotate(n)
			if n.right != nil && isRed(n.right.left) {
				flipColors(n)
			}
		}
	} else if isRed(n.right) && !isRed(n.left) {
		if isRed(n.right.left) && !isRed(n.right.right) {
			n.right = rightRotate(n.right)
		}
		if isRed(n.right.right) {
			n = leftRotate(n)
			if n.left != nil && isRed(n.left.right) {
				flipColors(n)
			}
		}
	} else if isRed(n.left) && isRed(n.right) {
		if isRed(n.left.left) || isRed(n.left.right) || isRed(n.right.left) || isRed(n.right.right) {
			flipColors(n)
		}
	}
	n.n = getSize(n)
	return n
}

func (t *Tree234) Remove(key interface{}) {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		if t.Compare(t.root.key, key) == 0 {
			t.root = nil
		}
		return
	}

	if !t.Contains(key) {
		return
	}

	if !isRed(t.root.left) && isRed(t.root.right) {
		t.root.color = Red
	}
	t.root = t.remove(t.root, key)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func (t *Tree234) Range(f func(n common.INode)) {
	common.PreOrder(t.root, f)
}

// Img 生成图
func (t *Tree234) Img(filename string) error {
	if filename == "" {
		filename = "tree234"
	}
	if t.GetSize() > 0 {
		t.root.BuildIndex()
		return common.PrintTree(t.root, filename)
	}
	return nil
}

func (t *Tree234) String() string {
	return common.PrePrint(t.root)
}
