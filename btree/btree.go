package btree

import (
	"sort"

	"github.com/skeletongo/datastructure/common"
)

type BTree struct {
	m       int // 阶数
	root    *node
	Compare func(a, b interface{}) int
}

// New 创建b树
// m 阶数
// compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func New(m int, compare func(a, b interface{}) int) *BTree {
	if m < 3 {
		m = 3
	}
	return &BTree{m: m, Compare: compare}
}

func (b *BTree) GetSize() int {
	if b.root == nil {
		return 0
	}
	return b.root.n
}

func (b *BTree) IsEmpty() bool {
	return b.GetSize() == 0
}

func (b *BTree) Rank() int {
	return b.m
}

func inOrder(n *node, list *[]interface{}) {
	if n == nil {
		return
	}
	c := len(n.children)
	if c > 0 {
		i := 0
		for ; i < len(n.values); i++ {
			inOrder(n.children[i], list)
			*list = append(*list, n.values[i])
		}
		inOrder(n.children[i], list)
	} else {
		for i := 0; i < len(n.values); i++ {
			*list = append(*list, n.values[i])
		}
	}
}

func (b *BTree) isBST() bool {
	arr := new([]interface{})
	inOrder(b.root, arr)
	for i := 1; i < len(*arr); i++ {
		if b.Compare((*arr)[i-1], (*arr)[i]) > 0 {
			return false
		}
	}
	return true
}

func (b *BTree) preOrder(n *node) bool {
	if n == nil {
		return true
	}
	if len(n.values) == b.m {
		return false
	}
	if len(n.children) > 0 && len(n.values)+1 != len(n.children) {
		return false
	}
	for _, v := range n.children {
		if b.preOrder(v) {
			continue
		} else {
			return false
		}
	}
	return true
}

func (b *BTree) isBtree() bool {
	return b.preOrder(b.root)
}

// add 向一个节点添加一个键值对并维护键值对顺序
func (b *BTree) add(n *node, value interface{}) {
	index := sort.Search(len(n.values), func(i int) bool {
		return b.Compare(n.values[i], value) > 0
	})
	if index == len(n.values) {
		n.values = append(n.values, value)
	} else {
		n.values = append(n.values, nil)
		copy(n.values[index+1:], n.values[index:])
		n.values[index] = value
	}
}

// find 在节点n中查询指定键值key所在位置
// i 默认-1，如果当前节点包含查询的键值返回键值在node.values的索引
// j 默认-1，如果当前节点不包含这个键值返回下一个查询的位置node.children的索引
func (b *BTree) find(n *node, value interface{}) (i, j int) {
	i, j = -1, -1
	// 二分搜索
	index := sort.Search(len(n.values), func(i int) bool {
		return b.Compare(n.values[i], value) >= 0
	})
	if index < len(n.values) && b.Compare(n.values[index], value) == 0 {
		return index, j
	}
	// 节点的子节点，要么全有要么全没有
	if len(n.children) == 0 {
		return
	}
	// 节点n的键值都比key小
	if index == len(n.values) {
		return i, len(n.children) - 1
	}
	// 返回下一个查询位置
	return i, index
}

// siftUp 节点分解,并向父节点融合
func (b *BTree) siftUp(n *node, k int) {
	mid := len(n.values) >> 1
	v := n.values[mid]

	// 分解成两个节点
	newNode := new(node)
	newNode.values = make([]interface{}, len(n.values)-mid-1)
	copy(newNode.values, n.values[mid+1:])
	n.values = n.values[:mid]
	if len(n.children) > 0 {
		newNode.children = make([]*node, len(newNode.values)+1)
		copy(newNode.children, n.children[mid+1:])
		n.children = n.children[:mid+1]
	}
	newNode.n = len(newNode.values)
	for _, v := range newNode.children {
		newNode.n += v.n
		v.parent = newNode
	}

	if k != -1 {
		// 有父节点，将一个键值向父节点融合
		newNode.parent = n.parent
		pNode := n.parent
		pNode.values = append(pNode.values, nil)
		copy(pNode.values[k+1:], pNode.values[k:])
		pNode.values[k] = v

		pNode.children = append(pNode.children, nil)
		copy(pNode.children[k+2:], pNode.children[k+1:])
		pNode.children[k+1] = newNode
	} else {
		// 没有父节点，即当前节点为根节点，创建新的根节点
		n.n = len(n.values)
		for _, v := range n.children {
			n.n += v.n
		}

		root := new(node)
		root.values = append(root.values, v)
		root.children = []*node{n, newNode}

		newNode.parent = root
		n.parent = root

		root.n = 1 + n.n + newNode.n
		b.root = root
	}
}

func (b *BTree) put(n *node, value interface{}, k int) {
	i, j := b.find(n, value)
	if i != -1 {
		return
	}
	if j == -1 {
		b.add(n, value)
	} else {
		b.put(n.children[j], value, j)
	}
	// 节点分解,节点融合
	if len(n.values) == b.m {
		b.siftUp(n, k)
	}
	// 维护键值数量
	n.n = len(n.values)
	for _, v := range n.children {
		n.n += v.n
	}
}

// Put 添加键值对
// 如果是已存在的键，对应的值会被覆盖
// 若该节点元素个数小于m-1，直接插入；
// 若该节点元素个数等于m-1，引起节点分裂；以该节点中间元素为分界，取中间元素（偶数个数，中间两个随机选取）插入到父节点中；
// 重复上面动作，直到所有节点符合B树的规则；最坏的情况一直分裂到根节点，生成新的根节点，高度增加1；
func (b *BTree) Put(value interface{}) {
	if b.root == nil {
		b.root = &node{
			n:      1,
			values: []interface{}{value},
		}
		return
	}
	b.put(b.root, value, -1)
}

func (b *BTree) contains(n *node, value interface{}) bool {
	i, j := b.find(n, value)
	if i != -1 {
		return true
	}
	if j == -1 {
		return false
	}
	return b.contains(n.children[j], value)
}

func (b *BTree) Contains(value interface{}) bool {
	if b.root == nil {
		return false
	}
	return b.contains(b.root, value)
}

func (b *BTree) get(n *node, value interface{}) interface{} {
	i, j := b.find(n, value)
	if i != -1 {
		return n.values[i]
	}
	if j == -1 {
		return nil
	}
	return b.get(n.children[j], value)
}

func (b *BTree) Get(value interface{}) interface{} {
	if b.root == nil {
		return nil
	}
	return b.get(b.root, value)
}

//todo RemoveMin
func (b *BTree) RemoveMin() {

}

//todo RemoveMax
func (b *BTree) RemoveMax() {

}

//todo Remove
func (b *BTree) Remove(value interface{}) {

}

// Img 生成图片
func (b *BTree) Img(filename string) error {
	if filename == "" {
		filename = "BTree"
	}
	if b.GetSize() > 0 {
		return common.BTreeSvg(b.root, filename)
	}
	return nil
}

func (b *BTree) String() string {
	return common.PrePrintTree(b.root)
}
