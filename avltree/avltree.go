package avltree

import (
	"dataStructure/common"
	"fmt"
	"math"
)

type node struct {
	height      int // 以当前节点为根的树的高度
	left, right *node
	key, value  interface{} // 存储映射的键值对
}

func newNode(key, value interface{}) *node {
	return &node{
		height: 1,
		key:    key,
		value:  value,
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

type AVLTree struct {
	size    int
	root    *node
	Compare func(a, b interface{}) int
}

// New 创建AVL树
// 参数 f 为自定义元素大小比较函数
// 大小比较函数 返回值：
// -1	表示	a<b
// 0	表示	a=b
// 1	表示	a>b
func New(f func(a, b interface{}) int) *AVLTree {
	return &AVLTree{Compare: f}
}

// GetSize 获取元素数量
func (a *AVLTree) GetSize() int {
	return a.size
}

// IsEmpty 是否为空
func (a *AVLTree) IsEmpty() bool {
	return a.size == 0
}

// IsBST 判断是不是二分搜索树
func (a *AVLTree) IsBST() bool {
	list := new([]interface{})
	a.inOrder(a.root, list)
	for i := 1; i < len(*list); i++ {
		if a.Compare((*list)[i-1], (*list)[i]) > 0 {
			return false
		}
	}
	return true
}

func (a *AVLTree) inOrder(n *node, list *[]interface{}) {
	if n == nil {
		return
	}
	a.inOrder(n.left, list)
	*list = append(*list, n.key)
	a.inOrder(n.right, list)
}

// IsBalanced 判断是不是平衡二叉树
func (a *AVLTree) IsBalanced() bool {
	return a.isBalanced(a.root)
}

func (a *AVLTree) isBalanced(n *node) bool {
	if n == nil {
		return true
	}
	balanceFactor := a.getBalanceFactory(n)
	if balanceFactor > 1 || balanceFactor < -1 {
		return false
	}
	return a.isBalanced(n.left) && a.isBalanced(n.right)
}

// getHeight 获取节点高度
func (a *AVLTree) getHeight(n *node) int {
	if n == nil {
		return 0
	}
	return n.height
}

// getBalanceFactory 获取节点的平衡因子
func (a *AVLTree) getBalanceFactory(n *node) int {
	if n == nil {
		return 0
	}
	return a.getHeight(n.left) - a.getHeight(n.right)
}

// findMinNode 寻找最小子节点
func (a *AVLTree) findMinNode(n *node) *node {
	cur := n
	for cur.left != nil {
		cur = n.left
	}
	return cur
}

// leftRotate 左旋转
// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func (a *AVLTree) leftRotate(n *node) *node {
	x := n.right
	t := x.left

	n.right = t
	x.left = n

	// 维护节点高度
	n.height = 1 + int(math.Max(float64(a.getHeight(n.left)), float64(a.getHeight(n.right))))
	x.height = 1 + int(math.Max(float64(a.getHeight(x.left)), float64(a.getHeight(x.right))))
	return x
}

// rightRotate 右旋转
// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func (a *AVLTree) rightRotate(n *node) *node {
	x := n.left
	t := x.right

	n.left = t
	x.right = n

	// 维护节点高度
	n.height = 1 + int(math.Max(float64(a.getHeight(n.left)), float64(a.getHeight(n.right))))
	x.height = 1 + int(math.Max(float64(a.getHeight(x.left)), float64(a.getHeight(x.right))))
	return x
}

// toBalance 维护节点平衡性
func (a *AVLTree) toBalance(n *node) *node {
	// 计算平衡因子
	balanceFactor := a.getBalanceFactory(n)
	// 维护平衡
	// LL
	if balanceFactor > 1 && a.getBalanceFactory(n.left) >= 0 {
		return a.rightRotate(n)
	}
	// RR
	if balanceFactor < -1 && a.getBalanceFactory(n.right) <= 0 {
		return a.leftRotate(n)
	}
	// LR
	if balanceFactor > 1 && a.getBalanceFactory(n.left) < 0 {
		n.left = a.leftRotate(n.left)
		return a.rightRotate(n)
	}
	// RL
	if balanceFactor < -1 && a.getBalanceFactory(n.right) > 0 {
		n.right = a.rightRotate(n.right)
		return a.leftRotate(n)
	}
	return n
}

// Add 添加新节点
func (a *AVLTree) Add(key, value interface{}) {
	a.root = a.add(a.root, key, value)
}

func (a *AVLTree) add(n *node, key, value interface{}) *node {
	if n == nil {
		a.size++
		return newNode(key, value)
	}

	res := a.Compare(n.key, key)
	if res == 0 {
		return n
	} else if res < 0 {
		n.right = a.add(n.right, key, value)
	} else {
		n.left = a.add(n.left, key, value)
	}

	// 路径回溯维护平衡性

	oldHeight := n.height
	// 更新height值
	n.height = 1 + int(math.Max(float64(a.getHeight(n.left)), float64(a.getHeight(n.right))))

	// 对于添加新节点的父节点来说如果高度没有变说明父节点原来就有一个子节点，另外有一个空节点，新节点就添加在这个空节点上
	// 这样才会导致父节点的高度没有改变，而且父节点的平衡因子一定是0
	// 因为新添加节点的父节点的高度不变，所以继续回溯路径上的节点的平衡因子也不会改变，平衡性也就不会被打破
	// **对于删除节点的操作此判断不正确，原因请看删除部分的代码注释**
	if n.height == oldHeight {
		return n
	}
	return a.toBalance(n)
}

func (a *AVLTree) contains(n *node, key interface{}) bool {
	if n == nil {
		return false
	}

	r := a.Compare(n.key, key)
	if r < 0 {
		return a.contains(n.right, key)
	}
	if r > 0 {
		return a.contains(n.left, key)
	}
	return true
}

// Contains 查询是否包含指定元素
func (a *AVLTree) Contains(key interface{}) bool {
	return a.contains(a.root, key)
}

// Remove 删除节点
func (a *AVLTree) Remove(key int) {
	a.root = a.remove(a.root, key)
}

func (a *AVLTree) remove(n *node, key interface{}) *node {
	if n == nil {
		return nil
	}

	var retNode *node
	res := a.Compare(n.key, key)
	if res > 0 {
		retNode = a.remove(n.left, key)
	} else if res < 0 {
		retNode = a.remove(n.right, key)
	} else {
		a.size--
		// 当前节点就是要删除的节点
		if n.left == nil {
			r := n.right
			n.right = nil
			retNode = r
		} else if n.right == nil {
			l := n.left
			n.left = nil
			retNode = l
		} else {
			// 左右都有子树
			// 用右子树中的最小值节点代替当前删除的节点
			min := a.findMinNode(n.right)
			min.right = a.remove(n.right, min.key)
			a.size++
			min.left = n.left
			n.left = nil
			n.right = nil
			retNode = min
		}
	}

	// 路径回溯维护平衡性

	if retNode == nil {
		return nil
	}

	// 更新height值
	retNode.height = 1 + int(math.Max(float64(a.getHeight(n.left)), float64(a.getHeight(n.right))))

	// 删除节点时如果回溯路径上的节点高度值没有改变也不能保证平衡因子就不会改变，也就不能保证节点的平衡性不会被打破
	// 如：
	//      x		将z节点删除后x节点的高度没有改变，但是平衡性已经被打破
	//     /  \
	//    y    z
	//   /
	//   m
	return a.toBalance(retNode)
}

func (a *AVLTree) set(n *node, key, value interface{}) *node {
	if n == nil {
		a.size++
		return newNode(key, value)
	}

	res := a.Compare(n.key, key)
	if res < 0 {
		n.right = a.set(n.right, key, value)
		return n
	}
	if res > 0 {
		n.left = a.set(n.left, key, value)
		return n
	}
	n.value = value
	return n
}

// Set 添加元素或修改元素对应的值
func (a *AVLTree) Set(key, value interface{}) {
	a.root = a.set(a.root, key, value)
}

func (a *AVLTree) get(n *node, key interface{}) *node {
	if n == nil {
		return nil
	}

	res := a.Compare(n.key, key)
	if res < 0 {
		return a.get(n.right, key)
	}
	if res > 0 {
		return a.get(n.left, key)
	}
	return n
}

// Get 获取元素值
func (a *AVLTree) Get(key interface{}) interface{} {
	n := a.get(a.root, key)
	if n == nil {
		return nil
	}
	return n.value
}

func (a *AVLTree) String() string {
	return common.PrePrint(a.root)
}
