package bst

import "dataStructure/array"

/*
 二分搜索树的查询,添加,删除操作的时间复杂度为O(h) 最差时间复杂度O(n)链表 最佳时间复杂O(logN)满二叉树

 满二叉树：除了最大层节点，其他节点都有两个子节点
 完全二叉树：按每层从左到右的位置添加新节点
 平衡二叉树：所有叶子节点所在的层数的差值的绝对值不能大于1
*/

// 二分搜索树
type BST struct {
	root *node
	size int
	f    func(a, b interface{}) int
}

// 获取节点总数
func (b *BST) GetSize() int {
	return b.size
}

// 是否为空树
func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// 添加新节点
func (b *BST) Add(e interface{}) {
	b.root = b.add(b.root, e)
}

func (b *BST) add(node *node, e interface{}) *node {
	if node == nil {
		b.size++
		return newNode(e)
	}

	if n := b.f(e, node.e); n == 0 {
		return node
	} else if n < 0 {
		node.left = b.add(node.left, e)
	} else {
		node.right = b.add(node.right, e)
	}
	return node
}

// 查询是否包含某节点
func (b *BST) Contains(e interface{}) bool {
	return b.contains(b.root, e)
}

func (b *BST) contains(node *node, e interface{}) bool {
	if node == nil {
		return false
	}

	if n := b.f(e, node.e); n == 0 {
		return true
	} else if n < 0 {
		return b.contains(node.left, e)
	} else {
		return b.contains(node.right, e)
	}
}

// 前序遍历
func (b *BST) PreOrder(f func(e interface{})) {
	b.preOrder(b.root, f)
}

func (b *BST) preOrder(node *node, f func(e interface{})) {
	if node == nil {
		return
	}

	f(node.e)
	b.preOrder(node.left, f)
	b.preOrder(node.right, f)
}

// 中序遍历
func (b *BST) InOrder(f func(e interface{})) {
	b.inOrder(b.root, f)
}

func (b *BST) inOrder(node *node, f func(e interface{})) {
	if node == nil {
		return
	}

	b.inOrder(node.left, f)
	f(node.e)
	b.inOrder(node.right, f)
}

// 后序遍历
func (b *BST) PostOrder(f func(e interface{})) {
	b.postOrder(b.root, f)
}

func (b *BST) postOrder(node *node, f func(e interface{})) {
	if node == nil {
		return
	}

	b.postOrder(node.left, f)
	b.postOrder(node.right, f)
	f(node.e)
}

// 深度优先遍历(中序遍历非递归形式)
func (b *BST) DFS(f func(e interface{})) {
	if b.root == nil {
		return
	}

	stack := array.Stack{}
	stack.Init()
	stack.Push(b.root)

	push := true // 压入标志

	for !stack.IsEmpty() {
		ele := stack.Peek()
		node := ele.(*node)

		if push { // 压入
			if node.left != nil { // 有左节点就将左节点压入栈
				stack.Push(node.left)
			} else { // 没有左节点就弹出这个节点，将此节点的右节点压入栈
				push = false
				stack.Pop()
				f(node.e)

				if node.right != nil {
					stack.Push(node.right)
					push = true
				}
			}
		} else { // 弹出
			stack.Pop()
			f(node.e)

			if node.right != nil {
				stack.Push(node.right)
				push = true
			}
		}
	}
}

// 广度优先遍历(层序遍历)
func (b *BST) BFS(f func(e interface{})) {
	if b.root == nil {
		return
	}

	queue := array.LoopQueue{}
	queue.Init()
	queue.Enqueue(b.root)

	for !queue.IsEmpty() {
		ele := queue.Dequeue()
		node := ele.(*node)
		f(node.e)
		// 子节点入队
		if node.left != nil {
			queue.Enqueue(node.left)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}
}

// 删除最大元素
func (b *BST) RemoveMax() interface{} {
	res := b.findMax(b.root)
	b.root = b.removeMax(b.root)
	return res.e
}

func (b *BST) removeMax(node *node) *node {
	if node == nil {
		return nil
	}

	if node.right == nil {
		// 右子树为空，删除当前节点，返回当前节点的左子树
		retNode := node.left
		node.left = nil
		b.size--
		return retNode
	}

	node.right = b.removeMax(node.right)
	return node
}

func (b *BST) findMax(node *node) *node {
	if node == nil {
		panic("no data")
	}
	cur := node
	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

// 删除最小元素
func (b *BST) RemoveMin() interface{} {
	res := b.findMin(b.root)
	b.root = b.removeMin(b.root)
	return res.e
}

func (b *BST) removeMin(node *node) *node {
	if node == nil {
		return nil
	}

	if node.left == nil {
		// 左子树为空，删除当前节点，返回当前节点的右子树
		retNode := node.right
		node.right = nil
		b.size--
		return retNode
	}

	node.left = b.removeMin(node.left)
	return node
}

func (b *BST) findMin(node *node) *node {
	if node == nil {
		panic("no data")
	}

	cur := node
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// 删除任意元素
func (b *BST) Remove(e interface{}) bool {
	sz := b.size
	b.root = b.remove(b.root, e)
	return sz > b.size
}

func (b *BST) remove(node *node, e interface{}) *node {
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

// 创建二分搜索树
// 参数 f 为自定义元素大小比较函数
// 大小比较函数 返回值：
// -1	表示	a<b
// 0	表示	a=b
// 1	表示	a>b
//func f(a,b interface{}) int { }
func NewBST(f func(a, b interface{}) int) *BST {
	return &BST{f: f}
}
