package dataStructure

import (
	"fmt"
	"strings"
)

/*
 线段树 更新和查询数据的时间复杂度是O(logN)
 线段树是一种平衡二叉树,特别适用于动态数据的查询和修改操作
*/

// 线段树
type SegmentTree struct {
	data []int
	tree []int
	f    func(a, b int) int
}

func (s *SegmentTree) GetSize() int {
	return len(s.data)
}

func (s *SegmentTree) Get(index int) int {
	if index < 0 || index >= len(s.data) {
		panic("param error")
	}
	return s.data[index]
}

func (s *SegmentTree) NewTree(arr []int, f func(a, b int) int) {
	s.f = f
	// 创建一个副本
	s.data = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		s.data[i] = arr[i]
	}
	// 开辟4倍元素个数的存储空间
	s.tree = make([]int, 4*len(arr))

	s.buildTree(0, 0, len(s.data)-1)
}

func (s *SegmentTree) leftChild(k int) int {
	return 2*k + 1
}

func (s *SegmentTree) rightChild(k int) int {
	return 2*k + 2
}

// 创建线段树
// 在treeIndex的位置创建表示区间[l...r]的线段树,具体表示用f表示
// treeIndex 表示 tree 中的每个节点的下标
// l,r 表示 data 中下标从l到r的数据下标
func (s *SegmentTree) buildTree(treeIndex, l, r int) {
	if l == r { // 递归终止条件
		s.tree[treeIndex] = s.data[l]
		return
	}
	leftTreeIndex := s.leftChild(treeIndex)
	rightTreeIndex := s.rightChild(treeIndex)
	mid := l + (r-l)/2 // 防止整数溢出
	s.buildTree(leftTreeIndex, l, mid)
	s.buildTree(rightTreeIndex, mid+1, r)
	// 融合两个线段树
	s.tree[treeIndex] = s.f(s.tree[leftTreeIndex], s.tree[rightTreeIndex])
}

// 返回区间[ql,qr]的值
func (s *SegmentTree) Query(ql, qr int) int {
	if ql < 0 || ql >= len(s.data) || qr < 0 || qr >= len(s.data) || ql > qr {
		panic("param error")
	}
	return s.query(0, 0, len(s.data)-1, ql, qr)
}

func (s *SegmentTree) query(treeIndex, l, r, ql, qr int) int {
	if l == ql && r == qr {
		return s.tree[treeIndex]
	}
	mid := l + (r-l)/2
	leftTreeIndex := s.leftChild(treeIndex)
	rightTreeIndex := s.rightChild(treeIndex)
	if qr <= mid {
		return s.query(leftTreeIndex, l, mid, ql, qr)
	}
	if ql >= mid+1 {
		return s.query(rightTreeIndex, mid+1, r, ql, qr)
	}
	return s.f(s.query(leftTreeIndex, l, mid, ql, mid), s.query(rightTreeIndex, mid+1, r, mid+1, qr))
}

// 修改线段树
func (s *SegmentTree) Set(index int, e int) {
	if index < 0 || index >= len(s.data) {
		panic("param error")
	}
	s.set(0, 0, len(s.data)-1, index, e)
}

func (s *SegmentTree) set(treeIndex, l, r, index, e int) {
	if l == r {
		s.tree[treeIndex] = e
		return
	}
	mid := l + (r-l)/2
	leftTreeIndex := s.leftChild(treeIndex)
	rightTreeIndex := s.rightChild(treeIndex)
	if index <= mid {
		s.set(leftTreeIndex, l, mid, index, e)
	} else {
		s.set(rightTreeIndex, mid+1, r, index, e)
	}
	// 修改叶子节点值后修改父节点的值
	s.tree[treeIndex] = s.f(s.tree[leftTreeIndex], s.tree[rightTreeIndex])
}

func (s *SegmentTree) String() string {
	buf := strings.Builder{}
	buf.WriteString("[")
	for k, v := range s.tree {
		buf.WriteString(fmt.Sprint(v))
		if k != len(s.tree)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
