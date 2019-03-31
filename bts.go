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
