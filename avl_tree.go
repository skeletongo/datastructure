package dataStructure

import "math"

type AVLTree struct {
	size int
	root *AVLNode
}

func (a *AVLTree) GetSize() int {
	return a.size
}

func (a *AVLTree) IsEmpty() bool {
	return a.size == 0
}

// 判断是否是二分搜索树
func (a *AVLTree) IsBST() bool {
	list := &[]int{}
	a.inOrder(a.root, list)
	for i := 1; i < len(*list); i++ {
		if (*list)[i-1] > (*list)[i] {
			return false
		}
	}
	return true
}

func (a *AVLTree) inOrder(node *AVLNode, list *[]int) {
	if node == nil {
		return
	}
	a.inOrder(node.Left, list)
	*list = append(*list, node.Key)
	a.inOrder(node.Right, list)
}

// 判断是否是平衡二叉树
func (a *AVLTree) IsBalanced() bool {
	return a.isBalanced(a.root)
}

func (a *AVLTree) isBalanced(node *AVLNode) bool {
	if node == nil {
		return true
	}
	balanceFactor := a.GetBalanceFactory(node)
	if balanceFactor > 1 {
		return false
	}
	return a.isBalanced(node.Left) && a.isBalanced(node.Right)
}

// 获取节点高度
func (a *AVLTree) GetHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// 获取节点的平衡因子
func (a *AVLTree) GetBalanceFactory(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return a.GetHeight(node.Left) - a.GetHeight(node.Right)
}

// 添加新节点
func (a *AVLTree) Add(key int) {
	a.root = a.add(a.root, key)
}

func (a *AVLTree) add(node *AVLNode, key int) *AVLNode {
	if node == nil {
		a.size++
		return NewAVLNode(key)
	}

	if node.Key == key {
		return node
	} else if node.Key < key {
		node.Right = a.add(node.Right, key)
	} else {
		node.Left = a.add(node.Left, key)
	}
	// 更新height值
	node.Height = 1 + int(math.Max(float64(a.GetHeight(node.Left)), float64(a.GetHeight(node.Right))))
	// 计算平衡因子
	balanceFactor := a.GetBalanceFactory(node)
	// 维护平衡
	// LL
	if balanceFactor > 1 && a.GetBalanceFactory(node.Left) >= 0 {
		return a.rightRotate(node)
	}
	// RR
	if balanceFactor < -1 && a.GetBalanceFactory(node.Right) <= 0 {
		return a.leftRotate(node)
	}
	// LR
	if balanceFactor > 1 && a.GetBalanceFactory(node.Left) < 0 {
		node.Left = a.leftRotate(node.Left)
		return a.rightRotate(node)
	}
	// RL
	if balanceFactor < -1 && a.GetBalanceFactory(node.Right) > 0 {
		node.Right = a.rightRotate(node.Right)
		return a.leftRotate(node)
	}
	return node
}

// 左旋转
func (a *AVLTree) leftRotate(node *AVLNode) *AVLNode {
	x := node.Right
	t := x.Left

	node.Right = t
	x.Left = node
	// 维护节点高度
	node.Height = 1 + int(math.Max(float64(a.GetHeight(node.Left)), float64(a.GetHeight(node.Right))))
	x.Height = 1 + int(math.Max(float64(a.GetHeight(x.Left)), float64(a.GetHeight(x.Right))))
	return x
}

// 右旋转
func (a *AVLTree) rightRotate(node *AVLNode) *AVLNode {
	x := node.Left
	t := x.Right

	node.Left = t
	x.Right = node
	// 维护节点高度
	node.Height = 1 + int(math.Max(float64(a.GetHeight(node.Left)), float64(a.GetHeight(node.Right))))
	x.Height = 1 + int(math.Max(float64(a.GetHeight(x.Left)), float64(a.GetHeight(x.Right))))
	return x
}

func (a *AVLTree) Remove(key int) {
	a.root = a.remove(a.root, key)
}

func (a *AVLTree) remove(node *AVLNode, key int) *AVLNode {
	if node == nil {
		return nil
	}

	var resNode *AVLNode
	if key < node.Key {
		resNode = a.remove(node.Left, key)
	} else if key > node.Key {
		resNode = a.remove(node.Right, key)
	} else {
		a.size--
		// 当前节点就是要删除的节点
		if node.Left == nil {
			tmp := node.Right
			node.Right = nil
			resNode = tmp
		} else if node.Right == nil {
			tmp := node.Left
			node.Left = nil
			resNode = tmp
		} else {
			// 左右都有子树
			// 用右子树中的最小值节点代替当前删除的节点
			min := a.findMinNode(node.Right)
			min.Right = a.remove(node.Right, min.Key)
			a.size++
			min.Left = node.Left
			node.Left = nil
			node.Right = nil
			resNode = min
		}
	}

	// 更新height值
	resNode.Height = 1 + int(math.Max(float64(a.GetHeight(resNode.Left)), float64(a.GetHeight(resNode.Right))))
	// 计算平衡因子
	balanceFactor := a.GetBalanceFactory(resNode)
	// 维护平衡
	// LL
	if balanceFactor > 1 && a.GetBalanceFactory(resNode.Left) >= 0 {
		return a.rightRotate(resNode)
	}
	// RR
	if balanceFactor < -1 && a.GetBalanceFactory(resNode.Right) <= 0 {
		return a.leftRotate(resNode)
	}
	// LR
	if balanceFactor > 1 && a.GetBalanceFactory(resNode.Left) < 0 {
		resNode.Left = a.leftRotate(resNode.Left)
		return a.rightRotate(resNode)
	}
	// RL
	if balanceFactor < -1 && a.GetBalanceFactory(resNode.Right) > 0 {
		resNode.Right = a.rightRotate(resNode.Right)
		return a.leftRotate(resNode)
	}

	return resNode
}

// 寻找最小子节点
func (a *AVLTree) findMinNode(node *AVLNode) *AVLNode {
	//if node == nil {
	//	panic("no min node")
	//}
	tmp := node
	for tmp.Left != nil {
		tmp = node.Left
	}
	return tmp
}
