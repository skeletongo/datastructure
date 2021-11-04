// Package tree23 左倾红黑树（2-3树）
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
	*list = append(*list, n.value)
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

func preOrder(n *node) bool {
	if n == nil {
		return true
	}

	if n.right != nil && isRed(n.right) {
		return false
	}
	if isRed(n) && isRed(n.left) {
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
func (t *Tree23) isBalanced() bool {
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

func (t *Tree23) put(n *node, value interface{}) *node {
	if n == nil {
		return newNode(value)
	}

	res := t.Compare(n.value, value)
	if res > 0 {
		n.left = t.put(n.left, value)
	} else if res < 0 {
		n.right = t.put(n.right, value)
	} else {
		n.value = value
		return n
	}

	// 路径回溯维护平衡性
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

func (t *Tree23) Put(value interface{}) {
	t.root = t.put(t.root, value)
	t.root.color = Black
}

func (t *Tree23) contains(n *node, value interface{}) bool {
	if n == nil {
		return false
	}

	res := t.Compare(n.value, value)
	if res > 0 {
		return t.contains(n.left, value)
	}
	if res < 0 {
		return t.contains(n.right, value)
	}
	return true
}

func (t *Tree23) Contains(value interface{}) bool {
	return t.contains(t.root, value)
}

func (t *Tree23) get(n *node, value interface{}) interface{} {
	if n == nil {
		return nil
	}

	res := t.Compare(n.value, value)
	if res > 0 {
		return t.get(n.left, value)
	}
	if res < 0 {
		return t.get(n.right, value)
	}
	return n.value
}

func (t *Tree23) Get(value interface{}) interface{} {
	return t.get(t.root, value)
}

// colorsFlip 颜色翻转
func colorsFlip(n *node) {
	n.color = Black
	n.left.color = Red
	n.right.color = Red
}

// n为红节点，并且左右子节点都是黑色，都不为空
func moveRedLeft(n *node) *node {
	// n由红色变黑色，子节点由黑色变红色，当前节点变成临时4节点
	// 此时如果右节点是3节点，从右节点借一个节点给左节点
	colorsFlip(n)
	if isRed(n.right.left) {
		n.right = rightRotate(n.right)
		// 这里可以不左旋转，也就是不借，不借也可以，这样在回溯维护平衡性的时候会比原来多旋转一次，但上面的右旋转必须有
		n = leftRotate(n)
	}
	return n
}

// removeMin 从3节点中删除最小值，所以传入的是3节点中的黑节点或3节点中的红节点
func (t *Tree23) removeMin(n *node) *node {
	// 左节点为空，当前节点就是最小节点
	if n.left == nil {
		return nil
	}

	// 两个性质
	// 1.在2-3树中2节点和3节点的子节点要么全有要么全没有，所以红色节点的子节点要么全有要么全没有都为空
	// 2.红色节点的子节点都是黑色

	// 根据方法的定义n有两种情况
	// n 是红色：由1可得左节点必定是黑色
	// n 是黑色：由方法定义可得左节点必定是红色（左节点还是2-3树中3节点中的节点所以递归调用此方法继续搜索最小节点）
	if !isRed(n.left) && !isRed(n.left.left) {
		// 因为当前节点的左节点是黑色节点，所以当前节点是红色节点
		// 并且由1可得右节点一定存在并且是黑节点
		// 因为左节点的左节点也是黑色，所以左节点在2-3树中是2节点，此时左节点需要向右节点借一个节点或者和父节点和右节点组成一个4节点
		n = moveRedLeft(n)
	}
	n.left = t.removeMin(n.left)

	// 路径回溯维护平衡性
	if isRed(n.right) && !isRed(n.left) {
		n = leftRotate(n)
	}
	// 删除最小值不会出现连续两个红色右节点的情况，因为默认红色节点在左侧；
	// 当前节点可能是个4节点，并且左节点的左节点是红节点，这样的话只需要分解4节点就可以；
	if isRed(n.left) && isRed(n.left.left) && !isRed(n.right) {
		n = rightRotate(n)
	}
	// 分解4节点
	if isRed(n.left) && isRed(n.right) {
		flipColors(n)
	}
	n.n = getSize(n)
	return n
}

func (t *Tree23) RemoveMin() {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		t.root = nil
		return
	}

	// 左节点是黑色，根节点也是黑色
	// 删除节点要在3节点中删除，所以根节点要变成红色
	if !isRed(t.root.left) {
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

	// 路径回溯维护平衡性
	if isRed(n.right) && !isRed(n.left) {
		n = leftRotate(n)
	}
	// 删除最大值会出现连续两个红色右节点的情况
	if isRed(n.left) && isRed(n.left.right) {
		n.left = leftRotate(n.left)
	}
	// 删除最大值会出现连续两个红色左节点的情况，因为默认红色节点在左侧，当寻找最大值的过程中颜色翻转时并且最大值就是右节点；
	// 当前节点可能是个4节点，并且左节点的左节点是红节点，这样的话只需要分解4节点就可以；
	if isRed(n.left) && isRed(n.left.left) && !isRed(n.right) {
		n = rightRotate(n)
	}
	// 分解4节点
	if isRed(n.left) && isRed(n.right) {
		flipColors(n)
	}
	n.n = getSize(n)
	return n
}

func (t *Tree23) RemoveMax() {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		t.root = nil
		return
	}

	if !isRed(t.root.left) {
		t.root.color = Red
	}
	t.root = t.removeMax(t.root)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func minNode(n *node) interface{} {
	cur := n
	for cur.left != nil {
		cur = cur.left
	}
	return cur.value
}

func (t *Tree23) remove(n *node, value interface{}) *node {
	res := t.Compare(n.value, value)
	if res > 0 {
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
		}
		n.left = t.remove(n.left, value)
	} else {
		if isRed(n.left) {
			n = rightRotate(n)
		}
		if t.Compare(n.value, value) == 0 && n.right == nil {
			return nil
		}
		if !isRed(n.right) && !isRed(n.right.left) {
			n = moveRedRight(n)
		}
		if t.Compare(n.value, value) == 0 {
			n.value = minNode(n.right)
			n.right = t.removeMin(n.right)
		} else {
			n.right = t.remove(n.right, value)
		}
	}

	if isRed(n.right) && !isRed(n.left) {
		n = leftRotate(n)
	}
	if isRed(n.left) && isRed(n.left.right) {
		n.left = leftRotate(n.left)
	}
	if isRed(n.left) && isRed(n.left.left) && !isRed(n.right) {
		n = rightRotate(n)
	}
	if isRed(n.left) && isRed(n.right) {
		flipColors(n)
	}
	n.n = getSize(n)
	return n
}

func (t *Tree23) Remove(value interface{}) {
	if t.root == nil {
		return
	}

	if t.root.n == 1 {
		if t.Compare(t.root.value, value) == 0 {
			t.root = nil
		}
		return
	}

	if !t.Contains(value) {
		return
	}

	if !isRed(t.root.left) {
		t.root.color = Red
	}
	t.root = t.remove(t.root, value)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func (t *Tree23) Range(f func(interface{})) {
	common.PreOrder(t.root, f)
}

// Img 生成图片
func (t *Tree23) Img(filename string) error {
	if filename == "" {
		filename = "tree23"
	}
	if t.GetSize() > 0 {
		return common.BSTSvg(t.root, filename)
	}
	return nil
}

func (t *Tree23) String() string {
	return common.PrePrintBST(t.root)
}
