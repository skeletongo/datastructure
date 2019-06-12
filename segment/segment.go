package segment

import (
	"fmt"
	"strings"
)

/*
 线段树 更新和查询数据的时间复杂度是O(logN)
 线段树是一种平衡二叉树,特别适用于动态数据的查询和修改操作
*/

// 线段树
type Tree struct {
	data []interface{}                      // 原数据
	tree []interface{}                      // 线段树
	f    func(a, b interface{}) interface{} // 节点融合方法
}

// 创建线段树
func New(arr []interface{}, f func(a, b interface{}) interface{}) *Tree {
	if len(arr) == 0 {
		panic("no data")
	}

	t := Tree{}
	t.f = f
	t.data = make([]interface{}, len(arr))
	for i := 0; i < len(arr); i++ {
		t.data[i] = arr[i]
	}
	t.tree = make([]interface{}, 4*len(arr))
	t.buildTree(0, 0, len(arr)-1)
	return &t
}

// 数据个数
func (t *Tree) GetSize() int {
	return len(t.data)
}

// 获取某个元素
func (t *Tree) Get(index int) interface{} {
	if index < 0 || index >= len(t.data) {
		panic("param error")
	}
	return t.data[index]
}

func (t *Tree) leftChild(k int) int {
	return 2*k + 1
}

func (t *Tree) rightChild(k int) int {
	return 2*k + 2
}

// 将data中的数据存储到tree中
// 在treeIndex的位置创建表示区间[l...r]的线段树,具体表示用f表示
// treeIndex 表示 tree 中的每个节点的下标
// l,r 表示 data 中下标从l到r的数据下标
func (t *Tree) buildTree(treeIndex int, l, r int) {
	if l == r {
		t.tree[treeIndex] = t.data[l]
		return
	}

	mid := l + (r-l)/2
	leftIndex := t.leftChild(treeIndex)
	rightIndex := t.rightChild(treeIndex)

	t.buildTree(leftIndex, l, mid)
	t.buildTree(rightIndex, mid+1, r)

	t.tree[treeIndex] = t.f(t.tree[leftIndex], t.tree[rightIndex])
}

// 查询
func (t *Tree) Query(ql, qr int) interface{} {
	// 参数检查
	if ql < 0 || ql >= t.GetSize() || qr < 0 || qr >= t.GetSize() || ql > qr {
		panic("param error")
	}
	return t.query(0, 0, t.GetSize()-1, ql, qr)
}

func (t *Tree) query(treeIndex, l, r, ql, qr int) interface{} {
	if l == ql && r == qr {
		return t.tree[treeIndex]
	}

	mid := l + (r-l)/2
	leftIndex := t.leftChild(treeIndex)
	rightIndex := t.rightChild(treeIndex)
	// 查询范围有三种情况：在左节点，在右节点，左右节点都有
	if qr <= mid {
		return t.query(leftIndex, l, mid, ql, qr)
	}
	if ql >= mid+1 {
		return t.query(rightIndex, mid+1, r, ql, qr)
	}
	// 左节点中的值和右节点中的值融合
	return t.f(t.query(leftIndex, l, mid, ql, mid), t.query(rightIndex, mid+1, r, mid+1, qr))
}

// 修改
func (t *Tree) Set(index int, data interface{}) {
	if index < 0 || index >= t.GetSize() {
		panic("param error")
	}
	t.set(0, 0, t.GetSize()-1, index, data)
}

func (t *Tree) set(treeIndex, l, r, index int, data interface{}) {
	if l == r {
		t.tree[treeIndex] = data
		return
	}

	mid := l + (r-l)/2
	leftIndex := t.leftChild(treeIndex)
	rightIndex := t.rightChild(treeIndex)
	// 分两种情况：要修改的元素在左节点，要修改的元素在右节点
	if index <= mid {
		t.set(leftIndex, l, mid, index, data)
	} else {
		t.set(rightIndex, mid+1, r, index, data)
	}
	// 修改父节点的值，维护线段树的性质
	t.tree[treeIndex] = t.f(t.tree[leftIndex], t.tree[rightIndex])
}

func (t *Tree) String() string {
	buf := strings.Builder{}
	buf.WriteString("[")
	for k, v := range t.tree {
		buf.WriteString(fmt.Sprint(v))
		if k != len(t.tree)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
