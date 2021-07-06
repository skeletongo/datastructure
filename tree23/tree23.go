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

func (t *Tree23) remove2(n *node, key interface{}) *node {
	res := t.Compare(n.key, key)
	if res > 0 {
		if n.left == nil {
			return n
		}
		if !isRed(n.left) && !isRed(n.left.left) {
			n = moveRedLeft(n)
		}
		n.left = t.remove2(n.left, key)
	} else if res < 0 {
		if isRed(n.left) {
			n = rightRotate(n)
		}
		if n.right == nil {
			return n
		}
		if !isRed(n.right) && !isRed(n.right.left) {
			n = moveRedRight(n)
		}
		n.right = t.remove2(n.right, key)
	} else {
		if n.left == nil && n.right == nil {
			return nil
		}
		if n.left != nil {
			n.key = maxKey(n.left)
			n.value = t.get(n.left, n.key)
			n.left = t.removeMax(n.left)
		} else {
			n.key = minKey(n.right)
			n.value = t.get(n.right, n.key)
			n.right = t.removeMin(n.right)
		}
	}

	if isRed(n.right) && !isRed(n.left) {
		n = leftRotate(n)
	}
	if isRed(n.left) && isRed(n.left.right) {
		n.left = leftRotate(n.left)
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

func (t *Tree23) remove(n *node, key interface{}) *node {
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

	if !t.Contains(key) {
		return
	}

	if !isRed(t.root.left) {
		t.root.color = Red
	}
	t.root = t.remove2(t.root, key)
	if !t.IsEmpty() {
		t.root.color = Black
	}
}

func (t *Tree23) Range(f func(n common.INode)) {
	common.PreOrder(t.root, f)
}

// Img 生成图
func (t *Tree23) Img(filename string) error {
	if filename == "" {
		filename = "tree23"
	}
	if t.GetSize() > 0 {
		t.root.BuildIndex()
		return common.PrintTree(t.root, filename)
	}
	return nil
}

func (t *Tree23) String() string {
	return common.PrePrint(t.root)
}
