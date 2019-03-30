package dataStructure

import "fmt"

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
	for prev.Left != nil || prev.Right != nil {
		if prev.Num == n {
			return
		}
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
func (b *Bst) PreOrder() {
	b.preOrder(b.node)
}

func (b *Bst) preOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Num)
	b.preOrder(node.Left)
	b.preOrder(node.Right)
}

// 前序遍历(非递归)
func (b *Bst) PreOrder2() {
	if b.node == nil {
		return
	}

	top := 0
	list := []*Node{b.node}
	for len(list) > 0 {
		e := list[top]
		list = list[:top]
		top--

		fmt.Print(e.Num)
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
func (b *Bst) InOrder() {
	b.inOrder(b.node)
}

func (b *Bst) inOrder(node *Node) {
	if node == nil {
		return
	}

	b.inOrder(node.Left)
	fmt.Println(node.Num)
	b.inOrder(node.Right)
}

// 中序遍历(非递归)
func (b *Bst) InOrder2() {
	if b.node == nil {
		return
	}

	top := 0
	list := []*Node{b.node}
	for len(list) > 0 {
		e := list[top]
		if e.Left != nil {
			list = append(list, e.Left)
			top++
		} else {
			list = list[:top]
			top--

			fmt.Println(e.Num)
			if e.Right != nil {
				list = append(list, e.Right)
				top++
			}
		}
	}
}

// 后序遍历(递归)
func (b *Bst) Postorder() {
	b.postOrder(b.node)
}

func (b *Bst) postOrder(node *Node) {
	if node == nil {
		return
	}

	b.postOrder(node.Left)
	b.postOrder(node.Right)
	fmt.Println(node.Num)
}

// 后序遍历(非递归)
func (b *Bst) PostOrder2() {
	if b.node == nil {
		return
	}

	var cur *Node
	top := 0
	list := []*Node{b.node}
	for len(list) > 0 {
		e := list[top]
		if e.Right == cur { // 相同弹出
			fmt.Println(e.Num)
			list = list[:top]
			top--
			cur = e
		} else { // 不同压入
			list = append(list, e.Right)
			top++
			if e.Left != nil {
				list = append(list, e.Left)
				top++
			}
		}
	}
}

// *层序遍历又名广度优先遍历
// 层序遍历
func (b *Bst) LevelOrder() {
	if b.node == nil {
		return
	}
	list := []*Node{b.node}
	for len(list) > 0 {
		e := list[0]
		list = list[:len(list)-1]
		fmt.Println(e.Num)
		if e.Left != nil {
			list = append(list, e.Left)
		}
		if e.Right != nil {
			list = append(list, e.Right)
		}
	}
}
