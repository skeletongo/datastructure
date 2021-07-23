package hashtable

import (
	"bytes"

	"github.com/skeletongo/datastructure/tree23"
)

type TableTree struct {
	size      int // 元素数量
	index     int // 取模用到的素数在 capacity 中的下标
	m         int
	tableTree []*tree23.Tree23
	Compare   func(a, b interface{}) int
}

func NewTableTree(compare func(a, b interface{}) int) *TableTree {
	ret := new(TableTree)
	ret.Compare = func(a, b interface{}) int {
		return compare(a.([]interface{})[0], b.([]interface{})[0])
	}
	ret.m = capacity[0]
	ret.tableTree = make([]*tree23.Tree23, ret.m)
	for i := 0; i < len(ret.tableTree); i++ {
		ret.tableTree[i] = tree23.New(ret.Compare)
	}
	return ret
}

func (t *TableTree) GetSize() int {
	return t.size
}

func (t *TableTree) IsEmpty() bool {
	return t.size == 0
}

func (t *TableTree) resize() {
	t.m = capacity[t.index]
	newTableTree := make([]*tree23.Tree23, t.m)
	for i := 0; i < len(newTableTree); i++ {
		newTableTree[i] = tree23.New(t.Compare)
	}
	for _, v := range t.tableTree {
		v.Range(func(val interface{}) {
			k := val.([]interface{})[0]
			newTableTree[hash(k, t.m)].Put(val)
		})
	}
	t.tableTree = newTableTree
}

func (t *TableTree) Set(key, value interface{}) {
	tree := t.tableTree[hash(key, t.m)]
	n := tree.GetSize()
	tree.Put([]interface{}{key, value})
	if tree.GetSize() > n {
		t.size++
		// 扩容
		if t.size > t.m*upperTol && t.index+1 < len(capacity) {
			t.index++
			t.resize()
		}
	}
}

func (t *TableTree) Remove(key interface{}) {
	tree := t.tableTree[hash(key, t.m)]
	n := tree.GetSize()
	tree.Remove([]interface{}{key})
	if tree.GetSize() < n {
		t.size--
		// 缩容
		if t.size < t.m*lowerTol && t.index-1 >= 0 {
			t.index--
			t.resize()
		}
	}
}

func (t *TableTree) Contains(key interface{}) bool {
	return t.tableTree[hash(key, t.m)].Contains([]interface{}{key})
}

func (t *TableTree) Get(key interface{}) interface{} {
	v := t.tableTree[hash(key, t.m)].Get([]interface{}{key})
	if v == nil {
		return nil
	}
	return v.([]interface{})[1]
}

func (t *TableTree) String() string {
	buf := bytes.Buffer{}
	for _, v := range t.tableTree {
		if v.GetSize() > 0 {
			buf.WriteString(v.String())
			buf.WriteString("\n")
		}
	}
	return buf.String()
}
