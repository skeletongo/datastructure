// Package uf 并查集，时间复杂度O(log*n), 效率低于O(1)高于O(logN)
// 适合频繁查询或建立连接的网络节点数据
package uf

import "fmt"

// IUnionFind 并查集
type IUnionFind interface {
	GetSize() int
	IsConnected(p, q int) bool
	Union(p, q int)
}

// IElement 数据节点
type IElement interface {
	// UniqueId 数据唯一索引 生成的索引要小于并查集初始化时的参数n的大小
	UniqueId() int
}

// FindParentFunc 查询父节点的方法
type FindParentFunc func(*UF, int) int

// UF 并查集
type UF struct {
	// 记录节点间的关系
	// 例如：
	// 有数据A和B; A.UniqueId() 返回 1; B.UniqueId() 返回2; 让A节点指向B节点，也就是A是B的子节点，B是A的父节点;
	// 则 data 由初始状态 {1:1,2:2} 变成 {1:2,2:2}; 下标是子节点唯一索引，下标对应的元素值是父节点的唯一索引;
	data []int

	// 优先级
	rank []int

	// 查询父节点的方法
	f FindParentFunc
}

// New 创建并查集
// n 并查集支持的最大数据量
func New(n int) *UF {
	uf := new(UF)
	for i := 0; i < n; i++ {
		uf.data = append(uf.data, i)
		uf.rank = append(uf.rank, 1)
	}
	uf.f = FindParentFuncLess
	return uf
}

// NewFunc 创建并查集
// n 并查集支持的最大数据量
func NewFunc(n int, f FindParentFunc) *UF {
	uf := new(UF)
	for i := 0; i < n; i++ {
		uf.data = append(uf.data, i)
		uf.rank = append(uf.rank, 1)
	}
	if f == nil {
		uf.f = FindParentFuncLess
	} else {
		uf.f = f
	}
	return uf
}

func (u *UF) Set(f FindParentFunc) {
	u.f = f
}

// GetSize 并查集支持的最大数据量
func (u *UF) GetSize() int {
	return len(u.data)
}

// IsConnected 两个节点是否已连接
func (u *UF) IsConnected(i, j int) bool {
	if i == j {
		return true
	}
	return u.f(u, i) == u.f(u, j)
}

// IsConnectedElements 两个节点是否已连接
func (u *UF) IsConnectedElements(a, b IElement) bool {
	return u.IsConnected(a.UniqueId(), b.UniqueId())
}

// Union 连接两个节点
func (u *UF) Union(i, j int) {
	i = u.f(u, i)
	j = u.f(u, j)
	if i == j {
		return
	}

	// 根据两个元素所在树的rank不同判断合并方向
	// 将rank低的集合合并到rank高的集合上
	if u.rank[i] < u.rank[j] {
		u.data[i] = j
	} else if u.rank[i] > u.rank[j] {
		u.data[j] = i
	} else {
		u.data[i] = j
		u.rank[j] += 1
	}
}

// UnionElements 连接两个节点
func (u *UF) UnionElements(a, b IElement) {
	u.Union(u.f(u, a.UniqueId()), u.f(u, b.UniqueId()))
}

// todo 优化打印方式
func (u *UF) String() string {
	return fmt.Sprint(u.data)
}

// FindParentFuncLess 寻找父节点，查询同时做少量路径压缩
var FindParentFuncLess = FindParentFunc(func(uf *UF, i int) int {
	if i < 0 || i >= len(uf.data) {
		panic("index out of Union-Find size")
	}
	for i != uf.data[i] {
		uf.data[i] = uf.data[uf.data[i]]
		i = uf.data[i]
	}
	return i
})

func findParentMore(uf *UF, i int) int {
	if i < 0 || i >= len(uf.data) {
		panic("index out of Union-Find size")
	}
	// 路径压缩，将子节点都直接指向根节点
	if i != uf.data[i] {
		// 将当前节点指向跟节点
		uf.data[i] = findParentMore(uf, uf.data[i])
	}
	return uf.data[i]
}

// FindParentFuncMore 寻找父节点，查询同时将查询节点及其所有上层节点都指向根节点
var FindParentFuncMore = FindParentFunc(findParentMore)
