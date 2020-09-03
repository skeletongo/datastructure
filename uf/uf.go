// 并查集，操作时间复杂度O(log*n), 效率低于O(1)高于O(logN)
// 时间复杂度：
package uf

// IUnionFind 并查集
type IUnionFind interface {
	Size() int
	IsConnected(p, q int) bool
	UnionElements(p, q int)
}

// IElement 数据节点
type IElement interface {
	// 获取数据集中单个数据的唯一索引（最大范围 [0, MaxInt-1]）
	// 生成的索引要小于并查集初始化时的参数n的大小
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

// TestNew 创建并查集
// n 并查集支持的最大数据量
func TestNew(n int, f FindParentFunc) *UF {
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

func (u *UF) Set(f FindParentFunc) {
	u.f = f
}

// Size 并查集支持的最大数据量
func (u *UF) Size() int {
	return len(u.data)
}

// IsConnected 两个数据节点是否已连接
func (u *UF) IsConnected(a, b IElement) bool {
	if a.UniqueId() == b.UniqueId() {
		return true
	}
	return u.f(u, a.UniqueId()) == u.f(u, b.UniqueId())
}

// UnionElements 连接两个数据节点
func (u *UF) UnionElements(a, b IElement) {
	i := u.f(u, a.UniqueId())
	j := u.f(u, b.UniqueId())

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

// FindParentFuncLess 寻找父节点，查询同时只将查询节点指向根节点
var FindParentFuncLess = FindParentFunc(func(u *UF, i int) int {
	if i < 0 || i >= len(u.data) {
		panic("index out of Union-Find size")
	}
	// 将当前节点指向根节点
	b := i
	for i != u.data[i] {
		i = u.data[i]
	}
	u.data[b] = i
	/*
		这种方式稍微慢一点
		for i != u.data[i] {
			u.data[i] = u.data[u.data[i]]
			i = u.data[i]
		}
	*/
	return i
})

func findParentMore(u *UF, i int) int {
	if i < 0 || i >= len(u.data) {
		panic("index out of Union-Find size")
	}
	// 路径压缩，将子节点都直接指向根节点
	if i != u.data[i] {
		// 将当前节点指向跟节点
		u.data[i] = findParentMore(u, u.data[i])
	}
	return u.data[i]
}

// FindParentFuncMore 寻找父节点，查询同时将查询节点及其所有上层节点都指向根节点
var FindParentFuncMore = FindParentFunc(findParentMore)
