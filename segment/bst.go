package segment

import "github.com/skeletongo/datastructure/common"

type node struct {
	left, right *node
	value       interface{}
}

func (n *node) GetLeftNode() common.INode {
	return n.left
}

func (n *node) GetRightNode() common.INode {
	return n.right
}

func (n *node) GetValue() interface{} {
	return n.value
}

type BSTSegment struct {
	data  []interface{}                      // 原数据
	root  *node                              // 线段树
	Merge func(a, b interface{}) interface{} // 节点融合方法
}

func NewBSTSegment(arr []interface{}, merge func(a, b interface{}) interface{}) *BSTSegment {
	if len(arr) == 0 {
		panic("NewBSTSegment: no data")
	}

	t := BSTSegment{}
	t.Merge = merge
	t.data = make([]interface{}, len(arr))
	copy(t.data, arr)
	t.root = t.buildTree(0, len(arr)-1)
	return &t
}

func (t *BSTSegment) GetSize() int {
	return len(t.data)
}

func (t *BSTSegment) Get(index int) interface{} {
	if index < 0 || index >= len(t.data) {
		panic("index out of bounds")
	}
	return t.data[index]
}

func (t *BSTSegment) buildTree(l, r int) *node {
	n := new(node)
	if l == r {
		n.value = t.data[l]
		return n
	}
	mid := l + (r-l)/2
	n.left = t.buildTree(l, mid)
	n.right = t.buildTree(mid+1, r)
	n.value = t.Merge(n.left.value, n.right.value)
	return n
}

func (t *BSTSegment) Query(ql, qr int) interface{} {
	// 参数检查
	if ql < 0 || ql >= t.GetSize() || qr < 0 || qr >= t.GetSize() || ql > qr {
		panic("Query: param error")
	}
	return t.query(t.root, 0, t.GetSize()-1, ql, qr)
}

func (t *BSTSegment) query(n *node, l, r, ql, qr int) interface{} {
	if l == ql && r == qr {
		return n.value
	}

	mid := l + (r-l)/2
	// 查询范围有三种情况：在左节点，在右节点，左右节点都有
	if qr <= mid {
		return t.query(n.left, l, mid, ql, qr)
	}
	if ql >= mid+1 {
		return t.query(n.right, mid+1, r, ql, qr)
	}
	// 左节点中的值和右节点中的值融合
	return t.Merge(t.query(n.left, l, mid, ql, mid), t.query(n.right, mid+1, r, mid+1, qr))
}

func (t *BSTSegment) Set(index int, data interface{}) {
	if index < 0 || index >= t.GetSize() {
		panic("Set: param error")
	}
	t.set(t.root, 0, t.GetSize()-1, index, data)
}

func (t *BSTSegment) set(n *node, l, r, index int, data interface{}) {
	if l == r {
		n.value = data
		return
	}

	mid := l + (r-l)/2
	// 分两种情况：要修改的元素在左节点，要修改的元素在右节点
	if index <= mid {
		t.set(n.left, l, mid, index, data)
	} else {
		t.set(n.right, mid+1, r, index, data)
	}
	// 修改父节点的值，维护线段树的性质
	n.value = t.Merge(n.left.value, n.right.value)
}

func (t *BSTSegment) String() string {
	return common.PrePrint(t.root)
}
