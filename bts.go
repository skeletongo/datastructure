package dataStructure

// 二分搜索树
type Bst struct {
	node *Node
	size int
}

func (b *Bst) GetSize() int {
	return b.size
}

func (b *Bst) IsEmpty() bool {
	return b.size == 0
}

// 插入元素(递归)
func (b *Bst) Add(n int) {
	b.node = b.add(b.node, n)
}

func (b *Bst) add(node *Node, n int) *Node {
	b.size++
	if node == nil {
		return NewNode(n)
	}
	if node.Num == n {
		return node
	}
	if n < node.Num {
		node.Left = b.add(node.Left, n)
	} else {
		node.Right = b.add(node.Right, n)
	}
	return node
}

// 插入元素(非递归)
func (b *Bst) Add2(n int) {
	if b.node == nil {
		b.node = NewNode(n)
		return
	}
	prev := b.node
	for prev.Num != n && ((n < prev.Num && prev.Left != nil) || (n > prev.Num && prev.Right != nil)) {
		if n < prev.Num {
			prev = prev.Left
		} else {
			prev = prev.Right
		}
	}
	if prev.Num == n {
		return
	}
	if n < prev.Num {
		prev.Left = NewNode(n)
	} else {
		prev.Right = NewNode(n)
	}
	b.size++
}

// 查询元素(递归)
func (b *Bst) Contains(n int) bool {
	return b.contains(b.node, n)
}

func (b *Bst) contains(node *Node, n int) bool {
	if node == nil {
		return false
	}
	if node.Num == n {
		return true
	}
	if n < node.Num {
		return b.contains(node.Left, n)
	} else {
		return b.contains(node.Right, n)
	}
}

// 查询元素(非递归)
func (b *Bst) Contains2(n int) bool {
	if b.node == nil {
		return false
	}
	cur := b.node
	for cur.Left != nil || cur.Right != nil {
		if cur.Num == n {
			return true
		}
		if n < cur.Num && cur.Left != nil {
			cur = cur.Left
		} else if n > cur.Num && cur.Right != nil {
			cur = cur.Right
		}
		return false
	}
	if cur.Num == n {
		return true
	}
	return false
}

// 前序遍历(递归)
func (b *Bst) PreOrder(f func(*Node)) {
	b.preOrder(b.node, f)
}

func (b *Bst) preOrder(node *Node, f func(*Node)) {
	if node == nil {
		return
	}

	f(node)
	b.preOrder(node.Left, f)
	b.preOrder(node.Right, f)
}

// 前序遍历(非递归)
func (b *Bst) PreOrder2(f func(*Node)) {
	if b.node == nil {
		return
	}

	top := 0
	list := []*Node{b.node}
	for len(list) > 0 {
		e := list[top]
		list = list[:top]
		top--

		f(e)
		if e.Right != nil {
			list = append(list, e.Right)
			top++
		}
		if e.Left != nil {
			list = append(list, e.Left)
			top++
		}
	}
}

// *中序遍历又名深度优先遍历
// 中序遍历(递归)
func (b *Bst) InOrder(f func(*Node)) {
	b.inOrder(b.node, f)
}

func (b *Bst) inOrder(node *Node, f func(*Node)) {
	if node == nil {
		return
	}

	b.inOrder(node.Left, f)
	f(node)
	b.inOrder(node.Right, f)
}

// 中序遍历(非递归)
func (b *Bst) InOrder2(f func(*Node)) {
	if b.node == nil {
		return
	}

	var s bool
	top := 0
	list := []*Node{b.node}
	for len(list) > 0 {
		e := list[top]
		if !s { // 压入
			if e.Left != nil {
				list = append(list, e.Left)
				top++
			} else {
				list = list[:top]
				top--
				f(e)
				s = true
				if e.Right != nil {
					list = append(list, e.Right)
					top++
					s = false
				}
			}
		} else { // 弹出
			list = list[:top]
			top--
			f(e)
			if e.Right != nil {
				list = append(list, e.Right)
				top++
				s = false
			}
		}
	}
}

// 后序遍历(递归)
func (b *Bst) PostOrder(f func(*Node)) {
	b.postOrder(b.node, f)
}

func (b *Bst) postOrder(node *Node, f func(*Node)) {
	if node == nil {
		return
	}

	b.postOrder(node.Left, f)
	b.postOrder(node.Right, f)
	f(node)
}

// 后序遍历(非递归)
func (b *Bst) PostOrder2(f func(*Node)) {
	if b.node == nil {
		return
	}

	var prev *Node
	top := 0
	list := []*Node{b.node}
	for len(list) > 0 {
		e := list[top]
		if (e.Left == nil && e.Right == nil) || (prev != nil && (e.Right == prev || e.Left == prev)) {
			list = list[:top]
			top--
			f(e)
			prev = e
		} else {
			if e.Right != nil {
				list = append(list, e.Right)
				top++
			}
			if e.Left != nil {
				list = append(list, e.Left)
				top++
			}
		}
	}
}

// *层序遍历又名广度优先遍历
// 层序遍历
func (b *Bst) LevelOrder(f func(node *Node)) {
	if b.node == nil {
		return
	}
	list := []*Node{b.node}
	for len(list) > 0 {
		e := list[0]
		list = list[1:]
		f(e)
		if e.Left != nil {
			list = append(list, e.Left)
		}
		if e.Right != nil {
			list = append(list, e.Right)
		}
	}
}

func (b *Bst) findMax() *Node {
	if b.node == nil {
		panic("no data")
	}
	cur := b.node
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

func (b *Bst) findMin() *Node {
	if b.node == nil {
		panic("no data")
	}
	cur := b.node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// 删除最大元素(递归)
func (b *Bst) RemoveMax() int {
	e := b.findMax()
	b.node = b.removeMax(b.node)
	return e.Num
}

func (b *Bst) removeMax(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Right == nil {
		n := node.Left
		node.Left = nil
		b.size--
		return n
	}
	node.Right = b.removeMax(node.Right)
	return node
}

// 删除最大元素(非递归)
func (b *Bst) RemoveMax2() int {
	if b.node == nil {
		panic("no data")
	}
	var prev *Node
	cur := b.node
	for cur.Right != nil {
		prev = cur
		cur = cur.Right
	}
	if prev == nil {
		if cur.Left == nil {
			b.node = nil
		} else {
			b.node = cur.Left
		}
	} else {
		if cur.Left == nil {
			prev.Right = nil
		} else {
			prev.Right = cur.Left
		}
	}
	b.size--
	return cur.Num
}

// 删除最小值(递归)
func (b *Bst) RemoveMin() int {
	e := b.findMin()
	b.node = b.removeMin(b.node)
	return e.Num
}

func (b *Bst) removeMin(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		n := node.Right
		node.Right = nil
		b.size--
		return n
	}
	node.Left = b.removeMin(node.Left)
	return node
}

// 删除最小元素(非递归)
func (b *Bst) RemoveMin2() int {
	if b.node == nil {
		panic("no data")
	}
	var prev *Node
	cur := b.node
	for cur.Left != nil {
		prev = cur
		cur = cur.Left
	}
	if prev == nil {
		if cur.Right == nil {
			b.node = nil
		} else {
			b.node = cur.Right
		}
	} else {
		if cur.Right == nil {
			prev.Left = nil
		} else {
			prev.Left = cur.Right
		}
	}
	b.size--
	return cur.Num
}

func (b *Bst) addNode(node1, node2 *Node) *Node {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 == nil && node2 != nil {
		return node2
	}
	if node1 != nil && node2 == nil {
		return node1
	}

	if node1.Num == node2.Num {
		return node1
	}
	if node2.Num < node1.Num {
		node1.Left = b.addNode(node1.Left, node2)
	} else {
		node1.Right = b.addNode(node1.Right, node2)
	}
	return node1
}

// 删除任一元素
func (b *Bst) Remove(n int) bool {
	if b.node == nil {
		return false
	}
	return b.remove(nil, b.node, n)
}

func (b *Bst) remove(prev, node *Node, n int) bool {
	if node == nil {
		return false
	}

	if node.Num == n {
		// 移除操作
		if prev == nil {
			if node.Right != nil {
				b.node = node.Right
				if node.Left != nil {
					b.node = b.addNode(b.node, node.Left)
				}
			} else if node.Left != nil {
				b.node = node.Left
			} else {
				b.node = nil
			}
		} else {
			if prev.Left == node {
				if node.Right != nil {
					prev.Left = node.Right
					if node.Left != nil {
						prev.Left = b.addNode(prev.Left, node.Left)
					}
				} else if node.Left != nil {
					prev.Left = node.Left
				} else {
					prev.Left = nil
				}
			} else {
				if node.Right != nil {
					prev.Right = node.Right
					if node.Left != nil {
						prev.Right = b.addNode(prev.Right, node.Left)
					}
				} else if node.Left != nil {
					prev.Right = node.Left
				} else {
					prev.Right = nil
				}
			}
		}
		b.size--
		return true
	}

	if n < node.Num {
		return b.remove(node, node.Left, n)
	} else {
		return b.remove(node, node.Right, n)
	}
}

// 删除任一元素
func (b *Bst) Remove2(n int) bool {
	s := b.size
	b.node = b.remove2(b.node, n)
	return b.size < s
}

func (b *Bst) remove2(node *Node, n int) *Node {
	if node == nil {
		return nil
	}
	if node.Num == n {
		b.size--
		if node.Right == nil {
			tmp := node.Left // 包括左节点为空的情况
			node.Left = nil
			return tmp
		}
		if node.Left == nil {
			tmp := node.Right // 包括右节点为空的情况
			node.Right = nil
			return tmp
		}
		// 待删除节点左右子数均不为空的情况
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		bst := &Bst{node: node.Right}
		minNode := bst.findMin()
		minNode.Right = bst.removeMin(bst.node)
		b.size++ // bst.removeMin(bst.node) 中减了一个，但是实际上没有减去，所以要加回来
		minNode.Left = node.Left
		node.Left = nil
		node.Right = nil
		return minNode
	}
	if n < node.Num {
		node.Left = b.remove2(node.Left, n)
	} else {
		node.Right = b.remove2(node.Right, n)
	}
	return node
}
