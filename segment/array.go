package segment

import (
	"github.com/skeletongo/dataStructure/common"
)

/*
 线段树 更新和查询数据的时间复杂度是O(logN)
 线段树是一种平衡二叉树,特别适用于动态数据的查询和修改操作
*/

type ArraySegment struct {
	data  []interface{}                      // 原数据
	tree  []interface{}                      // 线段树
	Merge func(a, b interface{}) interface{} // 节点融合方法
}

// NewArraySegment 创建线段树
// arr 原数据
// merge 数据融合方法
func NewArraySegment(arr []interface{}, merge func(a, b interface{}) interface{}) *ArraySegment {
	if len(arr) == 0 {
		panic("NewArraySegment: no data")
	}

	t := ArraySegment{}
	t.Merge = merge
	t.data = make([]interface{}, len(arr))
	copy(t.data, arr)
	t.tree = make([]interface{}, 4*len(arr))
	t.buildTree(0, 0, len(arr)-1)
	return &t
}

// GetSize 数据个数
func (t *ArraySegment) GetSize() int {
	return len(t.data)
}

// Get 获取某个元素
func (t *ArraySegment) Get(index int) interface{} {
	if index < 0 || index >= len(t.data) {
		panic("index out of bounds")
	}
	return t.data[index]
}

func leftChild(k int) int {
	return 2*k + 1
}

func rightChild(k int) int {
	return 2*k + 2
}

// buildTree 将data中的数据存储到tree中
// 在treeIndex的位置创建表示区间[l...r]的线段树
// treeIndex 表示 tree 中的每个节点的下标
// l,r 表示 data 中下标从l到r的数据下标
func (t *ArraySegment) buildTree(treeIndex, l, r int) {
	if l == r {
		t.tree[treeIndex] = t.data[l]
		return
	}

	mid := l + (r-l)/2
	leftIndex := leftChild(treeIndex)
	rightIndex := rightChild(treeIndex)

	t.buildTree(leftIndex, l, mid)
	t.buildTree(rightIndex, mid+1, r)

	t.tree[treeIndex] = t.Merge(t.tree[leftIndex], t.tree[rightIndex])
}

// Query 查询
func (t *ArraySegment) Query(ql, qr int) interface{} {
	// 参数检查
	if ql < 0 || ql >= t.GetSize() || qr < 0 || qr >= t.GetSize() || ql > qr {
		panic("Query: param error")
	}
	return t.query(0, 0, t.GetSize()-1, ql, qr)
}

// query 从下标为treeIndex的线段树中搜索区间[ql...qr]的融合值
func (t *ArraySegment) query(treeIndex, l, r, ql, qr int) interface{} {
	if l == ql && r == qr {
		return t.tree[treeIndex]
	}

	mid := l + (r-l)/2
	leftIndex := leftChild(treeIndex)
	rightIndex := rightChild(treeIndex)
	// 查询范围有三种情况：在左节点，在右节点，左右节点都有
	if qr <= mid {
		return t.query(leftIndex, l, mid, ql, qr)
	}
	if ql >= mid+1 {
		return t.query(rightIndex, mid+1, r, ql, qr)
	}
	// 左节点中的值和右节点中的值融合
	return t.Merge(t.query(leftIndex, l, mid, ql, mid), t.query(rightIndex, mid+1, r, mid+1, qr))
}

// Set 修改原数据
func (t *ArraySegment) Set(index int, data interface{}) {
	if index < 0 || index >= t.GetSize() {
		panic("Set: param error")
	}
	t.set(0, 0, t.GetSize()-1, index, data)
}

func (t *ArraySegment) set(treeIndex, l, r, index int, data interface{}) {
	if l == r {
		t.tree[treeIndex] = data
		return
	}

	mid := l + (r-l)/2
	leftIndex := leftChild(treeIndex)
	rightIndex := rightChild(treeIndex)
	// 分两种情况：要修改的元素在左节点，要修改的元素在右节点
	if index <= mid {
		t.set(leftIndex, l, mid, index, data)
	} else {
		t.set(rightIndex, mid+1, r, index, data)
	}
	// 修改父节点的值，维护线段树的性质
	t.tree[treeIndex] = t.Merge(t.tree[leftIndex], t.tree[rightIndex])
}

func (t *ArraySegment) String() string {
	return common.PrePrintBSTSlice(t.tree)
}
