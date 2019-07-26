package dataStructure

/*
 并查集，操作时间复杂度O(log*n), 效率低于O(1)高于O(logN)
*/

// 并查集
type UnionFind struct {
	e    []int // 节点元素只存放了父节点的数组下标
	rank []int // 优先级
}

func NewUF(size int) *UnionFind {
	u := UnionFind{}
	for i := 0; i < size; i++ {
		u.e = append(u.e, i)
		u.rank = append(u.rank, 1)
	}
	return &u
}

func (u *UnionFind) GetSize() int {
	return len(u.e)
}

func (u *UnionFind) find(p int) int {
	if p < 0 || p >= len(u.e) {
		panic("param error")
	}
	// 用递归方式进行路径压缩，将子节点都直接指向根节点
	if p != u.e[p] {
		u.e[p] = u.find(u.e[p])
	}
	return u.e[p]
}

func (u *UnionFind) IsConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UnionFind) UnionElements(p, q int) {
	p1 := u.find(p)
	q1 := u.find(q)

	if p1 == q1 {
		return
	}

	// 根据两个元素所在树的rank不同判断合并方向
	// 将rank低的集合合并到rank高的集合上
	if u.rank[p1] < u.rank[q1] {
		u.e[p1] = q1
	} else if u.rank[p1] > u.rank[q1] {
		u.e[q1] = p1
	} else {
		u.e[p1] = q1
		u.rank[q1] += 1
	}
}
